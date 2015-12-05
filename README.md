#fast-push: CLI Plugin

This plugin is developed with the idea of talking with <https://github.com/xiwenc/cf-fastpush-controller>. With the fastpush controller serving GET and PUT on /files resource on your cf deployed app, we can hastily swap source files, static resources (html, css, js). So you do not need to redo cf push whenever you have a change.

This controller is included in <https://github.com/mendix/cf-mendix-buildpack> only right now. We are going to do some work to make this more generic for everyone, the most straight forward one is forking existing buildpacks and implementing the controller but we will see..
