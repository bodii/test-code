// {} 标签内嵌表达式
var a = <div>
 { [1, 2].map(function(i) {return <span>{i / 2} </span>;}) }
 </div>