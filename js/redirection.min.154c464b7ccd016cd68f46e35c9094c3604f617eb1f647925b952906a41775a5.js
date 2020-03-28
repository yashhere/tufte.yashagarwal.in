function _httpsAlways(){if(location.protocol=='http:'&&(location.hostname!="127.0.0.1"&&location.hostname!="localhost"))
location.href=location.href.replace("http","https");}
function _notGitLab(domainTo=true){var x=location.hostname.search("gitlab.io")!=-1?true:false;if(x){location.href=domainTo;}}
function _notGitHub(domainTo=true){var x=location.hostname.search("github.io")!=-1?true:false;if(x){location.href=domainTo;}}
_notGitLab("https://yashagarwal.in");_notGitHub("https://yashagarwal.in");_httpsAlways();