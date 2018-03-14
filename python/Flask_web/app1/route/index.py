#!/usr/bin/env python3
# -*- coding:utf-8 -*-


from datetime import datetime
from flask import (
                render_template, 
                session,
                redirect,
                url_for,
                abort,
                flash,
                request,
                current_app,
            )
from flask.views import MethodView
from flask import Blueprint
from flask_wtf import FlaskForm
from wtforms import StringField, SubmitField
from wtforms.validators import Required
from library.decorators import (
                admin_required, 
                permission_required,
            )
from global_config.permission import Permission
from flask_login import login_required, current_user
from model.user import User
from model.role import Role
from model.post import Post
from library.userForm import EditProfileForm, EditProfileAdminForm
from library.postForm import PostForm


# 注册蓝图
index = Blueprint('index', __name__)

# 辅助类
class NameForm(FlaskForm):
    name = StringField('What is your name?', validators=[Required()])
    submit = SubmitField('Submit')

# 路由列表
"""
@index.route('/', methods=['GET', 'POST'])
@login_required
def home():
    title = 'Index'
    name = None
    form = NameForm()
    if form.validate_on_submit():
        name = form.name.data
        old_name = session.get('name')
        if old_name is not None and old_name != name:
            flash('Looks like you have changed your name!')
        session['name'] = name
        return redirect(url_for('index.home'))

    return render_template(
            'index/index.html', 
            current_time = datetime.now(),
            form = form,
            name = session.get('name'),
            title = title,
            )
"""

# 如果是管理员登录
@index.route('/admin')
@login_required
@admin_required
def for_admins_only():
    return 'For administrators!'

# 如果是主持人登录
@index.route('/moderator')
@login_required
@permission_required(Permission.MODERATE_COMMENTS)
def for_moderators_only():
    return 'For comment moderators!'


''' 
为了避免每次调用render_template()时都多添加一个模板参数。
可以使用上下文处理器。上下文处理器能让变量在所有模板中全局访问。 
'''
@index.app_context_processor
def inject_permissions():
    return dict(Permission=Permission)


# 用户资料页
@index.route('/user/<username>')
def user(username):
    title = 'User profile page'
    user = User.query.filter_by(username=username).first()
    if user is None:
        abort(404)
    posts = user.posts.order_by(Post.timestamp.desc()).all()
    return render_template(
                    'index/user.html', 
                    user=user,
                    title=title,
                    posts=posts,
                )

# 资料编辑页
@index.route('/edit-profile', methods=['GET', 'POST'])
@login_required
def edit_profile():
    form = EditProfileForm()
    if form.validate_on_submit():
        current_user.name = form.name.data
        current_user.about_me = form.about_me.data
        db.session.add(current_user)
        flash('Your profile has been update.')
        return redirect(url_for('.user', username=current_user.username))
    form.name.data = current_user.name
    form.location.data = current_user.location
    form.about_me.data = current_user.about_me
    return render_template('index/edit_profile.html', form=form, title='edit_profile')


# 管理员资料编辑页
@index.route('/edit-profile/<int:id>', methods=['GET', 'POST'])
@login_required
@admin_required
def edit_profile_admin(id):
    user = User.query.get_or_404(id)
    form = EditProfileAdminForm(user=user)
    if form.validate_on_submit():
        user.email = form.email.data
        user.username = form.email.data
        user.confirmed = form.confirmed.data
        user.role = Role.query.get(form.role.data)
        user.name = form.name.data
        user.location = form.location.data
        user.about_me = form.about_me.data
        db.session.add(user)
        flash('The profile has been updated.')
        return redirect(url_for('.user', username=user.username))
    form.email.data = user.email
    form.username.data = user.username
    form.confirmed.data = user.confirmed
    form.role.data = user.role_id
    form.name.data = user.name
    form.location.data = user.location
    form.about_me.data = user.about_me
    return render_template('index/edit_profile.html', form=form, user=user, title='edit_profile')


@index.route('/', method=['GET', 'POST'])
def home():
    form = PostForm()
    if current_user.can(Permission.WRITE_ARTICLES) and \
            form.validate_on_submit():
        post = Post(
            body=form.body.data, 
            author=current_user._get_current_object()
            )
        db.session.add(post)
        return redirect(url_for('.index'))
    page = request.args.get('page', 1, type=int)
    pagination = Post.query.order_by(
        Post.timestamp.desc()
        ).paginate(
            page, 
            per_page=current_app.config['FLASKY_POSTS_PER_PAGE'] or 20,
            error_out=False
        )
    posts = pagination.items
    # posts = Post.query.order_by(Post.timestamp.desc()).all()
    return render_template(
            'index/index.html', 
            form=form, 
            posts=posts,
            pagination=pagination,
            )


# 文章专页
@index.route('/post/<int:id>')
def post(id):
    post = Post.query.get_or_404(id)
    return render_template('post.html', posts=[post])


# 编辑博客文章页
@index.route('/post_edit/<int:id>', methods=['GET', 'POST'])
@login_required
def edit(id):
    post = Post.query.get_or_404(id)
    if currnet_user != post.author and \
        not current_user.can(Permission.ADMINISTER):
        abort(403)
    form = PostForm()
    if form.validate_on_submit():
        psot.body = form.body.data
        db.session.add(post)
        flash('The post has been updated.')
        return redirect(url_for('post', id=post.id))
    form.body.data = post.body
    return render_template('edit_post.html', form=form)


# 关注功能
@index.route('/follow/<username>')
@login_required
@permission_required(Permission.FOLLOW)
def follow(username):
    user = User.query.filter_by(username=username).first()
    if user is None:
        flash('Invalid user.')
        return redirect(url_for('.home'))
    if current_user.is_following(user):
        flash('You are alread following this user.')
        return redirect(url_for('.user', username=username))
    current_user.follow(user)
    flash('You are now following %s.' % username)
    return redirect(url_for('.user', username=username))


# 取消关注功能
@index.route('/unfollow/<username>')
@login_required
@permission_required(Permission.FOLLOW)
def unfollow(username):
    user = User.query.filter_by(username=username).first()
    if user is None:
        flash('Invalid user.')
        return redirect(url_for('.home'))
    if current_user.is_following(user) is None:
        flahs('You have not followed this user.')
        return redirect(url_for('.home'))
    current_user.unfollow(user)
    flash('You are now unfollowing %s.' % username)
    return redirect(url_for('.user', username=username))

# 关注者关系列表页
@index.route('/followers/<username>')
def followers(username):
    user = User.query.filter_by(username=username).first()
    if user is None:
        flash('Invalid user.')
        return redirect(url_for('.home'))
    page = request.args.get('page', 1, type=int)
    pagination = user.followers.print_usage(
            page,
            per_page=current_app.config['FLASKY_FOLLOWERS_PER_PAGE'],
            error_out=False,
        )
    follows = [
                {
                    'user': item.follower, 
                    'timestamp': item.timestamp
                } for item in pagination.itmes
            ]

    return render_template(
                'followers.html', 
                user=user,
                title='Followers of',
                endpoint='.followers',
                pagination=pagination,
                follows=follows,
                )