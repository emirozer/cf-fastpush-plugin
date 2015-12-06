#fast-push: CLI Plugin

This plugin is developed with the idea of talking with <https://github.com/xiwenc/cf-fastpush-controller>. With the fastpush controller serving GET and PUT on /files resource on your cf deployed app, we can hastily swap source files, static resources (html, css, js). So you do not need to redo cf push whenever you have a change.

This controller is included in <https://github.com/mendix/cf-mendix-buildpack> only right now. We are going to do some work to make this more generic for everyone, the most straight forward one is forking existing buildpacks and implementing the controller but we will see..

### Installation

The latest binaries for this plugin will be provided in the bin folder of the repository. You can also cross compile with the build.sh script in bin folder.

Run **cf install-plugin BINARY_FILENAME** to install a plugin. Replace **BINARY_FILENAME** with the path to and name of the binary file.

###Intended Usage

Given you have an app named exampleApp, and assuming you are in its directory on your workstation, made changes and you want to push. You should do a dry run first:

    cf fast-push exampleApp --dry
    cf fast-push exampleApp
