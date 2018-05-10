from django.conf.urls import url
from django.urls import path, reverse
from . import views

urlpatterns = [
    # index
    url(r'^$', views.index, name='index'),
    # ex: /5/
    url(r'^(?P<question_id>[0-9]+)/$', views.detail, name='detail'),
    # ex: /5/results/
    url(r'^(?P<question_id>[0-9]+)/results/$', views.results, name='results'),
    # ex: /5/vote/
    url(r'^(?P<question_id>[0-9]+)/vote/$', views.vote, name='vote'),
]