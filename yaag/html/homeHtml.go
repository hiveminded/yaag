package html

const HomeHtmlString = `<html>
<title> API Documentation </title>
<body> 

	BaseLink <a href= {{.BaseLink}} > {{.BaseLink}} </a>

	<p> CurrentPath = {{.CurrentPath}}

	<p> MethodType = {{.MethodType}} </p>

	<p> <H4> Request Headers </H4> </p>
	{{ range $key, $value := .RequestHeaders }}
   <li><strong>{{ $key }}</strong>: {{ $value }}</li>
    {{ end }}
	

</body>
</html>`
