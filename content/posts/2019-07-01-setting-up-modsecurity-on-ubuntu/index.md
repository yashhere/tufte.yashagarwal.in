+++
title = "Setting Up ModSecurity on Ubuntu"
author = ["Yash Agarwal"]
date = 2019-07-01T18:20:18+05:30
images = []
tags = ["Linux"]
categories = ["Technical"]
draft = false
+++

Recently, I am experimenting with Web Application Firewalls a lot. ModSecurity is one of them. It is the most famous and useful open-source Web Application Firewall (WAF) in existence. It is supported by various web servers such as Apache, Nginx, and IIS.

The job of ModSecurity is to sit in front of the application web server and check the incoming requests and outgoing responses to filter out malicious content. It does so by the use of powerful and complex regular expressions. ModSecurity uses a rule language for its rules. The rule language has variables and operators defined to aid in the process of parsing HTTP requests.

ModSecurity, in itself, cannot block or allow requests. It is just a rule engine. It requires rules to operate appropriately. That's where its sister project, Core Rule Set (CRS), comes into the picture. CRS is a rule set developed to be used with ModSecurity. It has been in active development for several years now and is very mature. Together, ModSecurity and CRS form a formidable defense against the widespread web application attacks.

Now that you know, what a WAF is, let's proceed to install ModSecurity on Ubuntu. I will be compiling ModSecurity's latest version on Ubuntu 18.04. We will also configure ModSecurity to use Core Rule Set.

## Installing Dependencies
ModSecurity requires some dependencies to work correctly. Let's install them -

First, upgrade the Ubuntu system.

{{< highlight bash >}}
sudo apt-get -y update
sudo apt-get -y upgrade
{{< /highlight >}}

Now install the dependencies.

{{< highlight bash >}}
sudo apt-get -y install git libtool dh-autoreconf pkgconf gawk libcurl4-gnutls-dev libexpat1-dev libpcre3-dev libssl-dev libxml2-dev libyajl-dev zlibc zlib1g-dev libxml2 libpcre++-dev libxml2-dev libgeoip-dev liblmdb-dev lua5.2-dev iputils-ping locales apache2 apache2-dev ca-certificates wget
{{< /highlight >}}

*Optional*: clean up the Ubuntu caches.

{{< highlight bash >}}
sudo apt-get clean && sudo rm -rf /var/lib/apt/lists/*
{{< /highlight >}}

Install `SSDeep` as well (as done [here](https://github.com/CRS-support/modsecurity-docker/blob/v3/apache-apache/Dockerfile))

{{< highlight bash >}}
cd ~
git clone https://github.com/ssdeep-project/ssdeep
cd ssdeep
./bootstrap
./configure
make
sudo make install
{{< /highlight >}}

## Compiling ModSecurity
Let's clone ModSecurity from Github.

{{< highlight bash >}}
cd ~
git clone -b v3/master --single-branch https://github.com/SpiderLabs/ModSecurity
cd ModSecurity
git submodule init
git submodule update
./build.sh
./configure
make                # takes ~8 minutes on AWS t2.micro
sudo make install
{{< /highlight >}}

## Compiling ModSecurity-apache connector
To configure it with Apache, we will require ModSecurity-apache connector. Let's install that as well.

{{< highlight bash >}}
cd ~
git clone https://github.com/SpiderLabs/ModSecurity-apache
cd ModSecurity-apache
./autogen.sh
./configure --with-libmodsecurity=/usr/local/modsecurity
make
sudo make install
{{< /highlight >}}

## Setting up CRS rules
Now, let's download CRS rule set as well.

{{< highlight bash >}}
cd ~
git clone -b v3.2/dev https://github.com/SpiderLabs/owasp-modsecurity-crs
sudo mv owasp-modsecurity-crs/ /usr/local/
{{< /highlight >}}

Rename CRS configuration file - 

{{< highlight bash >}}
sudo mv /usr/local/owasp-modsecurity-crs/crs-setup.conf.example /usr/local/owasp-modsecurity-crs/crs-setup.conf
{{< /highlight >}}

## Setting up ModSecurity
Now, we need to create a file in the Apache modules directory, so that Apache can know, how to activate ModSecurity.

Create `/etc/apache2/mods-enabled/security3.conf` file and paste the following contents -

{{< highlight bash >}}
LoadModule security3_module /usr/lib/apache2/modules/mod_security3.so
modsecurity on
modsecurity_rules_file '/etc/apache2/modsec/main.conf'
{{< /highlight >}}

As you can see, the last line in the above code block reference a file `main.conf` in a folder `modsec`. This folder will not be present by default. We need to create that.

{{< highlight bash >}}
sudo mkdir -p /etc/apache2/modsec
{{< /highlight >}}

Setup ModSecurity configuration file -

{{< highlight bash >}}
# enables Unicode support in ModSecurity
sudo wget -P /etc/apache2/modsec/ https://raw.githubusercontent.com/SpiderLabs/ModSecurity/v3/master/unicode.mapping

sudo wget -P /etc/apache2/modsec/ https://raw.githubusercontent.com/SpiderLabs/ModSecurity/v3/master/modsecurity.conf-recommended
sudo mv /etc/apache2/modsec/modsecurity.conf-recommended /etc/apache2/modsec/modsecurity.conf
{{< /highlight >}}

Change the SecRuleEngine directive in the configuration to change from the default "detection only" mode to actively dropping malicious traffic.

{{< highlight bash >}}
sudo sed -i 's/SecRuleEngine DetectionOnly/SecRuleEngine On/' /etc/apache2/modsec/modsecurity.conf
{{< /highlight >}}

Change the location of `modsec_audit.log` file to Apache log directory.

{{< highlight bash >}}
sudo sed -i 's/SecAuditLog \/var\/log\/modsec_audit.log/SecAuditLog \/var\/log\/apache2\/modsec_audit.log/' /etc/apache2/modsec/modsecurity.conf
{{< /highlight >}}

To configure ModSecurity to use CRS rule set, put the following text in `/etc/apache2/modsec/main.conf` file.

{{< highlight bash >}}
Include "/etc/apache2/modsec/modsecurity.conf"
Include "/usr/local/owasp-modsecurity-crs/crs-setup.conf"
Include "/usr/local/owasp-modsecurity-crs/rules/*.conf"
{{< /highlight >}}

Also enable some Apache modules for better functioning of ModSecurity.

{{< highlight bash >}}
sudo a2enmod unique_id headers rewrite actions dav dav_fs
{{< /highlight >}}

Now restart the Apache server

{{< highlight bash >}}
sudo systemctl restart apache2
{{< /highlight >}}

## Fixing some common issues
Sometimes, I had encountered errors when ModSecurity was not able to append logs to its log file. I figured out that ModSecurity did not have enough permissions to write that file. We can fix this issue quickly.

First, test if you really have this issue or not.

{{< highlight bash >}}
curl 'http://localhost/?q="><script>alert(1)</script>'
<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
<html><head>
<title>403 Forbidden</title>
</head><body>
<h1>Forbidden</h1>
<p>You dont have permission to access / on this server.<br /></p>
<hr>
<address>Apache/2.4.29 (Ubuntu) Server at localhost Port 80</address>
</body></html>
{{< /highlight >}}

Now go to Apache log directory and check the contents of `modsec_audit.log` file.

{{< highlight bash >}}
cd /var/log/apache2
tail modsec_audit.log
{{< /highlight >}}

You should see the following content -

{{< highlight bash >}}
---0LzdyETA---A--
[01/Jul/2019:14:42:41 +0000] 156199216179.666171 127.0.0.1 41824 ip-xxx-xx-xx-xx.ap-south-1.compute.internal 80
---0LzdyETA---B--
GET /?q="><script>alert(1)</script> HTTP/1.1
Host: localhost
User-Agent: curl/7.58.0
Accept: */*

---TqjMwy7h---D--

---TqjMwy7h---F--
HTTP/1.1 403

---TqjMwy7h---H--
ModSecurity: Warning. detected XSS using libinjection. [file "/usr/local/owasp-modsecurity-crs/rules/REQUEST-941-APPLICATION-ATTACK-XSS.conf"] [line "37"] [id "941100"] [rev ""] [msg "XSS Attack Detected via libinjection"] [data "Matched Data: XSS data found within ARGS:q: "><script>alert(1)</script>"] [severity "2"] [ver "OWASP_CRS/3.1.0"] [maturity "0"] [accuracy "0"] [tag "application-multi"] [tag "language-multi"] [tag "platform-multi"] [tag "attack-xss"] [tag "OWASP_CRS/WEB_ATTACK/XSS"] [tag "WASCTC/WASC-8"] [tag "WASCTC/WASC-22"] [tag "OWASP_TOP_10/A3"] [tag "OWASP_AppSensor/IE1"] [tag "CAPEC-242"] [hostname "localhost"] [uri "/"] [unique_id "156198848361.198287"] [ref "v8,27t:utf8toUnicode,t:urlDecodeUni,t:htmlEntityDecode,t:jsDecode,t:cssDecode,t:removeNulls"]
....
....

---TqjMwy7h---I--

---TqjMwy7h---J--

---TqjMwy7h---Z--
{{< /highlight >}}

If you do not see the following content, and the file is empty or it does not exist, then ModSecurity was not able to open this file for writing. Use the following fix -

{{< highlight bash >}}
# find out the user, Apache is running as
apache_user="$(ps -ef | egrep '(httpd|apache2|apache)' | grep -v `whoami` | grep -v root | head -n1 | awk '{print $1}')"
{{< /highlight >}}

Add this user to `adm` group which owns the Apache logs directory in Ubuntu.

{{< highlight bash >}}
sudo usermod -G adm www-data
{{< /highlight >}}

Now, change the owner of Apache log directory to `apache_user`.

{{< highlight bash >}}
sudo chown -R $apache_user:$apache_user /var/log/apache2/*
{{< /highlight >}}

Now, ModSecurity should be able to append logs to the file `modsec_audit.log`.

## *Bonus*: Enabling JSON logs
**Note:** Honestly speaking, I was not able to make it work every time. I do not know what is the issue, but it works with some of the installations, and with some of the installations, it just doesn't log anything to the `audit` directory. If anyone has managed to make it work consistently, please let me know.

**Edit (13/07/2020):** The JSON logging works fine. The issue was that ModSecurity did not have permission to create subdirectories in the Apache log directory. I suppose it is something related to SELinux. However, a simple solution is to add the user under which the Apache process runs to the `adm` group. It might not be the right solution security-wise. However, from a quick remediation point of view, it works. Please let me know if you identify any better solution to fix the problem.

Anyway, if you are like me, who do not like the default ModSecurity log format, ModSecurity provides an option to generate logs in JSON format as well. To enable JSON support, the YAJL library should be installed. We already installed this package when we were installing dependencies, so our ModSecurity setup is compiled with JSON support. Let us now configure ModSecurity to generate JSON logs.

Open the `/etc/apache2/modsec/modsecurity.conf` file and find the following lines -

{{< highlight bash >}}
SecAuditLogType           Serial
SecAuditLog               /var/log/modsec_audit.log
{{< /highlight >}}

Once you have found the following lines, replace these lines with the following lines

{{< highlight bash >}}
SecAuditLogFormat         JSON
SecAuditLogType           Parallel
SecAuditLog               /var/log/apache2/modsec_audit.log
SecAuditLogStorageDir     /var/log/apache2/audit/

SecAuditLogFileMode       0644
SecAuditLogDirMode        0755
{{< /highlight >}}

Restart Apache server

{{< highlight bash >}}
sudo systemctl restart apache2
{{< /highlight >}}

Now, go to `/var/log/apache2/` directory and create `audit` folder.

{{< highlight bash >}}
sudo usermod -G adm $apache_user

cd /var/log/apache2
sudo mkdir audit

# make `apache_user` owner of this directory as well...
sudo chown -R $apache_user:$apache_user /var/log/apache2/audit
{{< /highlight >}}

Now, ModSecurity should be able to generate JSON logs in this directory. ModSecurity generates logs in the following format -

{{< highlight conf>}}
ubuntu@server:/var/log/apache2$ tree audit
audit
└── 20190701
    ├── 20190701-1132
    │   ├── 20190701-113225-156196094515.868593
    │   └── 20190701-113226-156196094691.154769
    ├── 20190701-1211
    │   ├── 20190701-121122-156196328239.048942
    │   └── 20190701-121122-156196328243.018882

    ....
    ....
{{< /highlight >}}

Now, your site should be relatively more secure than before.

## A warning, though
CRS is known to generate a lot of false-positive when enabled completely. We have not touched CRS paranoia levels. By default, it is set to paranoia level 1, which is known to produce false positives rarely, but still, as a measure of precaution, monitor your site's traffic for some time, and then decide if you need to disable some of the CRS rules or not.
