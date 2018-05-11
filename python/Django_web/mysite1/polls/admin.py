from django.contrib import admin
from .models import Question, Choice
# Register your models here.


# 选项
# class ChoiceInline(admin.StackedInline):
#     model = Choice
#     extra = 3
# v 2.0
class ChoiceInline(admin.TabularInline): # TabularInline 以紧凑的方式显示
    model = Choice
    extra = 3


# 问题
class QuestionAdmin(admin.ModelAdmin):
    # fields = ['pub_date', 'question_text']
    # v 2.0
    fieldsets = [
        (None,               {'fields': ['question_text']}),
        ('Date information', {'fields': ['pub_date'], 'classes': ['collapse']}),
    ]
    list_display = ('question_text', 'pub_date', 'was_published_recently') # 问题列表页的显示字段
    inlines = [ChoiceInline]
    search_fields = ['question_text'] # 添加搜索引并指定字段，以like来查询
    # fieldsets 元组中的第一个元素是字段集的标题

admin.site.register(Question, QuestionAdmin)

