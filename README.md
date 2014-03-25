Skate
=====

Command line utility to use Zurb's Ink Inliner (http://zurb.com/ink/inliner.php) for HTML emails.


### Usage

Given you have an html file:

``` html
<html>
  <head>
    <style>
      body{
        margin:20px;
      }
    </style>
  </head>
  <body></body>
</html>
```

Just pass a html file path as argument to Skate:

``` zsh
› skate index.html
```
And you'll get an inlined version:

``` html
<!-- Inliner Build Version 4380b7741bb759d6cb997545f3add21ad48f010b -->
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN" "http://www.w3.org/TR/REC-html40/loose.dtd">
<html>
<head></head>
<body style="margin: 20px;"></body>
</html>
```

Skate tries to be a good Unix citizen. That means that you can use [stdin](https://en.wikipedia.org/wiki/Standard_streams#Standard_input_.28stdin.29) and [stdout](https://en.wikipedia.org/wiki/Standard_streams#Standard_output_.28stdout.29) to pipe in/out data with the Inliner:

``` zsh
› date | skate
```

``` html
<!-- Inliner Build Version 4380b7741bb759d6cb997545f3add21ad48f010b -->
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN" "http://www.w3.org/TR/REC-html40/loose.dtd">
<html><body><p>martes, 25 de marzo de 2014, 13:44:41 GMT
</p></body></html>
```

Or in a more realistic example:

``` zsh
› skate index.html > inlined.html
```

### Why inline styles?

Inline styles is a good technique to achieve consistent rendering when sending email newsletters to difference email client. As mentioned by Zurb's guys:

> * Popular email clients like Gmail strip out CSS in the \<style\> tag.
> * This is the best way to guarantee your email works properly on the compatible clients listed in the docs.


### Disclaimer

This is _literally_ my first piece of code in [Go](golang.org). Please, be understanding.

