+++
title = "Setting up Pyhton Development Environments"
date = "2016-10-09"
categories = ["python", "development"]
tags = ["pip","python", "development"]
type = "posts"
+++
Recently I was searching for python projects on Github for contribution. Every single project I found, had a thing common among them. In every project's contribution guide, it was asked to setup the virtual environment for project. What the hack is this virtual environment and how does it work?

As a beginner to open source projects, the problem I faced in the beginning was how to set up the development environments for the projects I was looking at. I searched internet, I found some articles but they were not complete. So I decided to write my own guide, which will be useful for me in future also.

Python uses <code>pip</code> for package management.

### Installing pip
<code>pip</code> depends on setuptools library, which is in official Ubuntu repositories. To install it for python2:
<pre>
    sudo apt-get install python-setuptools
</pre>
Then install <code>pip</code> using
<pre>
    sudo apt-get install python-pip
</pre>

and for python3
<pre>
    sudo apt-get install python3-setuptools
</pre>
Then install <code>pip</code> using
<pre>
    sudo apt-get install python3-pip
</pre>

It should install <code>pip</code> on your system for both python versions. <code>pip</code> is very easy to use. It will take care of every single package you may require for your project.

#### Installing a package using pip
<pre>
    #it will search and install [package]
    pip install [package]
    pip install django
</pre>
If you are using python3, then don't forget to use <code>pip3</code>.

<code>pip</code> can be used to install a specific version of package also.
<pre>
    #it will search and install [package] with [version]
    pip install [package]==[version]
    pip install django==1.6.5
</pre>

#### Uninstalling a package using pip
<pre>
    #it will search and uninstall [package]
    pip uninstall [package]
    pip uninstall django
</pre>

#### upgrading a package using pip
<pre>
    #it will upgrade [package] to latest version
    pip install --upgrade [package]
    pip install --upgrade django
</pre>

#### Creating list of all packages with pip
This is one of most used and most useful feature of <code>pip</code>. It allows you to make a list of all the dependencies of your project.
<pre>
    #it will output the file to current directory
    pip freeze > [file_name.txt]
</pre>

All these commands above will install the packages globally. But that's not what is desired. <code>virtualenv</code> comes to our rescue here.

#### Virtualenv
<code>virtualenv</code> solves a very specific problem; it allows multiple python projects that have different and often conflicting dependencies, to coexist on same system.

<code>virtualenv</code> solves this problem by creating differet isolated development environments for your projects. An environment is a folder which contain everything, your project needs to work properly.

#### Installing virtualenv
By default, if you install <code>virtualenv</code> using <code>pip</code>, it will use system's default python to create virtual environments. To overcome this problem, we will install <code>virtualenv</code> using ubuntu package manager.
<pre>
    sudo apt-get install python-virtualenv
</pre>

#### Installing virtualenvwrapper
<code>virtualenvwrapper</code> provides some set of commands which makes working with virtual environments much more easy.

To install it,
<pre>
    sudo pip install virtualenvwrapper
</pre>

<code>pip</code>, <code>virtualenv</code> and <code>virtualenvwrapper</code> are the only packages which you will need to install globally. All other per project packages will be installed in respective virtual environments.

<code>virtualenvwrapper</code> also places all your virtual environments in one place. This makes working with projects very easy.

Now open your <code>.bashrc</code> and add these two lines to the end.
<pre>
    # All your projects will be saved in python-dev folder
    export PROJECT_HOME=~/python-dev
    # ~/python-dev/virtualenvs will contains python interpreters for each project.
    export WORKON_HOME=~/python-dev/virtualenvs
    # source the virtualenvwrapper script
    source /usr/local/bin/virtualenvwrapper.sh
</pre>

You can change "python-dev" to any name you wish. Your virtual environments will be created at that location.

Now restart your terminal to source the <code>.bashrc</code> or use
<pre>
    source .bashrc
</pre>

####Basic Usage
Create a virtual environment
<pre>
    mkvirtualenv myproject
</pre>

This will create myproject folder in the python-dev directory. To activate a project
<pre>
    workon myproject
</pre>

Alternatively you can create project using <code>mkproject</code> command. It will create virtual environment as well as a project directory in the <code>$PROJECT_HOME</code>, which is cd -ed into when you workon myproject.

Don't forget to deactivate current project when you switch between different projects.

To deactivate a project:
<pre>
    deactivate
</pre>

To delete a virtual environment:
<pre>
    rmvirtualenv myproject
</pre>

List all environments
<pre>
    lsvirtualenv
</pre>

it will also list all virtual environments:
<pre>
    workon
</pre>

[Full List of virtualenvwrapper commands](https://virtualenvwrapper.readthedocs.io/en/latest/command_ref.html)

<code>virtualenvwrapper</code> also provides tab-completion feature which is very handy when you have a lot of projects to work with.

That's it. Hope you liked the post...:)

