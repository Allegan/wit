# wit
Wit (Git with a W) is a command line tool for fetching and uploading pages to MediaWiki based sites. It allows for pages to be fetched, written to a file, or written to STDOUT in order to feed the content into other shell-stlye applications for processing. 

Wit also supports the ability to take a file, or text from STDOUT, and upload it to MediaWiki based sites either to create new pages or to overwrite existing ones. The latter will count as an edit instead of a page creation. Anonymous editing is supported but disabled by default and requires a flag to enable. Anonymous editing is advised against as it displays the IP address that you are connecting from to all of the users of the wiki. In addition, editing anonymously has more restrictions placed on it, depending on the configuration of the site, than registered editing.

## Useage
```
~$ wit -f -t "Main Page" -s "en.wikipedia.org" -o "Main Page.txt"
~$ wit -e -t "Main Page" -s "en.wikipedia.org" -i "Main Page.txt"
Username: 
Password: 
```