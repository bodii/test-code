### 通用视图
1. ListView   显示一个对象列表
2. DetailView 显示一个特定类型对象的详细信息页面

：example
```
\# urls.py
\# 通用视图 -- 首页
path('', views.IndexView.as_view(), name='index'), 
\# 通过视图 -- 详情
path('<int:pk>/', views.DetailView.as_view(), name='detail')

\# views.py
from django.views import generic

class IndexView(generic.ListView): 
    template_name = 'polls/index.html'
    context_object_name = 'latest_question_list'

    def get_queryset(self):
        """Retrun the last five published questions."""
        return Question.objects.order_by('-pub_date')[:5]

class DetailView(generic.DetailView):
    model = Question
    template_name = 'polls/detail.html'
```

每个通用视图需要知道它将作用于哪个模型。这由model属性提供
DetailView期望从URL中捕获名为pk的主键值，所以我们为通用视图把question_id改成pk

默认情况下，通用视图DetailView使用一个叫做<app name>/<model name>_detail.html的模板。在
在例子中，它将使用“polls/question_detail.html”模板。template_name属性是用来告诉Djange使用一个
指定的模板名字，而不是自动生成的默认名字。我们也为results列表视图指定了template_name
这确保results视图和detail视图在渲染时具有不同的外观，即使它们在后台都是同一个DetailView。
类似地，ListView使用一个叫做<app name>/<model name>_list.html的默认模板；我们使用
template_name来告诉ListView使用我们创建的已经存在的"polls/index.html"模板

在之前的教程中，提供模板文件时都带有一个包含question和latest_question_list变量的context。对于DetailView, question变量会自动提供--因为我们使用Django的模型（Question）,Django能够为context变量决定一个合适的名字。然而对于ListView,自动生成
context变量是question_list。为了覆盖这个行为，我们提供context_object_name属性，表
示我们想使用latest_question_list。作为一种替换方案，可以改变你的模板来匹配新的context
变量--这是一种更便捷的方法，告诉Django使用你想使用的变量名。
