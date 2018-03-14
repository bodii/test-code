#!/usr/bin/env python3
# -*- coding:utf-8 -*-

from flask import render_template
from flask import Blueprint

error = Blueprint('error', __name__)


@error.app_errorhandler(403)
def resources_not_available(e):
    return render_template('403.html'), 403

@error.app_errorhandler(404)
def page_not_found(e):
    return render_template('404.html'), 404

@error.app_errorhandler(500)
def internal_server_error(e):
    return render_template('500.html'), 500