from django.shortcuts import render, get_object_or_404
from django.http import HttpResponse, Http404, HttpResponseRedirect
from django.urls import reverse
from django.template import loader
from django.db.models import F
from django.views import generic
from django.utils import timezone

from .models import Question, Choice


# Create your views here.

# def index(request):
#     latest_question_list = Question.objects.order_by('-pub_date')[:5]
#     # template = loader.get_template('polls/index.html')
#     context = {
#         'latest_question_list': latest_question_list,
#     }
#     # return HttpResponse(template.render(context, request))
#     return render(
#         request, 
#         'polls/index.html', 
#         context, 
#         # content_type='application/xhtml+xml', # 以 xml显示
#         content_type='text/html;charset=utf-8', # 以网页显示
#         )

# 通用视图
class IndexView(generic.ListView): 
    template_name = 'polls/index.html'
    context_object_name = 'latest_question_list'

    # v 2.0
    # def get_queryset(self):
    #     """Retrun the last five published questions."""
    #     return Question.objects.order_by('-pub_date')[:5]
    # v 3.0
    def get_queryset(self):
        """
        Return the last five published questions (not including those
        set to be published in the future).
        """
        return Question.objects.filter(
            pub_date__lte = timezone.now()
        ).order_by('-pub_date')[:5]

# def detail(request, question_id):
#     # return HttpResponse("You're looking at question %s." % question_id)
#     '''try:
#         question = Question.objects.get(pk=question_id)
#     except Question.DoesNotExist:
#         raise Http404('Question does not exist')'''
#     question = get_object_or_404(Question, pk=question_id)
#     return render(request, 'polls/detail.html', {'question': question})
class DetailView(generic.DetailView):
    model = Question
    template_name = 'polls/detail.html'

    # v 3.0
    def get_queryset(self):
        """
        Excludes any questions that aren't published yet.
        """
        return Question.objects.filter(pub_date__lte=timezone.now())

# def results(request, question_id):
#     question = get_object_or_404(Question, pk=question_id)
#     return render(request, 'polls/results.html', {'question': question})
class ResultsView(generic.DetailView):
    model = Question
    template_name = 'polls/results.html'

    def get_queryset(self):
        return Question.objects.filter(pub_date__lte=timezone.now())


def vote(request, question_id):
    # return HttpResponse("You're voting on question %s." % question_id)
    question = get_object_or_404(Question, pk=question_id)
    try:
        selected_choice = question.choice_set.get(pk=request.POST['choice'])
    except (KeyError, Choice.DoesNotExist):
        return render(request, 'polls/detail.html', {
            'question': question,
            'error_message': "You didn't select a choice.",
        })
    else:
        selected_choice.votes = F('votes') + 1
        selected_choice.save()
        return HttpResponseRedirect(reverse('polls:results', args=(question.id,)))
