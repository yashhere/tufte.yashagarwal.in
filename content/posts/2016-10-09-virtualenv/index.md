+++
title = "Setting up Python Development Environments"
author = ["Yash Agarwal"]
date = 2016-10-09
images = []
tags = ["Pip", "Python", "Virtual Environments"]
categories = ["Technical"]
draft = false
+++

Recently I was searching for Python projects on Github for contribution. Every single project I found, had a thing common among them. In every project's contribution guide, it was asked to set up the virtual environment for the project. What the heck is this virtual environment and how does it work?

As a beginner to open source projects, the problem I faced, in the beginning, was how to set up the development environments for the projects I was looking at. I searched the Internet, I found some articles, but they were not complete. So I decided to write this guide, which will be useful for me in future also.

Python uses `pip` for package management.

## Installing pip
`pip` depends on setuptools library, which is in official Ubuntu repositories. To install it for python2 -

{{< highlight bash >}}
sudo apt-get install python-setuptools
{{< /highlight >}}

Then install `pip` using -
{{< highlight bash >}}
sudo apt-get install python-pip
{{< /highlight >}}

and for python3 -
{{< highlight bash >}}
sudo apt-get install python3-setuptools
{{< /highlight >}}

Then install `pip` using -
{{< highlight bash >}}
sudo apt-get install python3-pip
{{< /highlight >}}

It should install `pip` on your system for both python versions. `pip` is very easy to use. It will take care of every single package you may require for your project.

### Installing a package using pip

{{< highlight bash >}}
# it will search and install [package]
pip install [package]
pip install django
{{< /highlight >}}
If you are using python3, then don't forget to use `pip3`.

`pip` can be used to install a specific version of package also.
{{< highlight bash >}}
# it will search and install [package] with [version]
pip install [package]==[version]
pip install django==1.6.5
{{< /highlight >}}

### Uninstalling a package using pip
{{< highlight bash >}}
# it will search and uninstall [package]
pip uninstall [package]
pip uninstall django
{{< /highlight >}}

### upgrading a package using pip
{{< highlight bash >}}
# it will upgrade [package] to latest version
pip install --upgrade [package]
pip install --upgrade django
{{< /highlight >}}

### Creating list of all packages with pip
It is one of most used and most useful feature of `pip`. It allows you to make a list of all the dependencies of your project.
{{< highlight bash >}}
# it will output the file to current directory
pip freeze > [file_name.txt]
{{< /highlight >}}

All these commands above will install the packages globally. But that's not what is desired. `virtualenv` comes to our rescue here.

## Virtualenv
`virtualenv` solves a very particular problem; it allows multiple python projects that have different and often conflicting dependencies, to coexist on the same system.

`virtualenv` solves this problem by creating different isolated development environments for your projects. An environment is a folder which contains everything; your project needs to work properly.

### Installing virtualenv
By default, if you install `virtualenv` using `pip`, it will use system's default python to create virtual environments. To overcome this problem, we will install `virtualenv` using ubuntu package manager.
{{< highlight bash >}}
sudo apt-get install python-virtualenv
{{< /highlight >}}

### Installing virtualenvwrapper
`virtualenvwrapper` provides some set of commands which makes working with virtual environments much easier.

To install it -
{{< highlight bash >}}
sudo pip install virtualenvwrapper
{{< /highlight >}}

`pip`, `virtualenv` and `virtualenvwrapper` are the only packages which you will need to install globally. All other per project packages will be installed in respective virtual environments.

`virtualenvwrapper` also places all your virtual environments in one place. It makes working with projects very easy.

Now open your `.bashrc` and add these two lines to the end -
{{< highlight bash >}}
# All your projects will be saved in python-dev folder
export PROJECT_HOME=~/python-dev

# ~/python-dev/virtualenvs will contains python interpreters for each project.
export WORKON_HOME=~/python-dev/virtualenvs

# source the virtualenvwrapper script
source /usr/local/bin/virtualenvwrapper.sh
{{< /highlight >}}

You can change `python-dev` to any name you wish. Your virtual environments will be created at that location.

Now restart your terminal to source the `.bashrc` or use -
{{< highlight bash >}}
source .bashrc
{{< /highlight >}}

### Basic Usage
Create a virtual environment -
{{< highlight bash >}}
mkvirtualenv myproject
{{< /highlight >}}

It will create `myproject` folder in the python-dev directory. To activate this project -
{{< highlight bash >}}
workon myproject
{{< /highlight >}}

Alternatively you can create project using `mkproject` command. It will create a virtual environment as well as a project directory in the `$PROJECT_HOME`, which is `cd`-ed into when you `workon` myproject.

Don't forget to deactivate current project when you switch between different projects.

To deactivate a project -
{{< highlight bash >}}
deactivate
{{< /highlight >}}

To delete a virtual environment -
{{< highlight bash >}}
rmvirtualenv myproject
{{< /highlight >}}

List all environments -
{{< highlight bash >}}
lsvirtualenv
{{< /highlight >}}

it will also list all virtual environments -
{{< highlight bash >}}
workon
{{< /highlight >}}

Please refer to virtualenvwrapper documentation for [full list of virtualenvwrapper commands](https://virtualenvwrapper.readthedocs.io/en/latest/command_ref.html).

virtualenvwrapper also provides the tab-completion feature which is very handy when you have a lot of projects to work with.

That's it. Hope you liked the post. :smile:

