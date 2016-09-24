A Simple, small, powerful OCR web server

its use to Convert Image to text

to install it in your server

1- you nede golang, leptonica-dev, tesseract-ocr, libtesseract-dev , gcc, g++ git
 
2- Download the program from this link:

https://notabug.org/alimiracle/Uruk-Ocr-Server/releases 
then extract it ( in home folder recomended ).

then go to program folder by typing:

  cd ocr-server
  
then type:

chmod +x install.sh

./install.sh as root

to config the programme 

open /etc/ocrconfig/server.conf as root

sudo nano /etc/ocrconfig/server.conf
replace http://myip.com/  with your server host
if your host is local
Type http://localhost/
replace 8080 with your server port You want your irc server listen to it

then save the file

To add new language
open /etc/ocrconfig/lang.conf

nano /etc/ocrconfig/lang.conf
and add the new language to the list

for  example
to add The Arabic language
the config file  looks like

{
"Lang":"ar"
}

and you can add More than language

for  example

to add English and Arabic

the config file  looks like

{
"Lang":"eng ar"
}

to run the program

type:

sudo /bin/ocr-server

And if you want to run it as  service

type:
systemctl enable ocr-server


have fun and be free

ali miracle
