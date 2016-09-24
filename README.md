irc log recorder is irc bot to  Watching Everything that happens in the irc channel 
and record it in log.
and its have web server To enable users to access the logs from Web Browser and telnet.

To install the program:
1- you nede install golang to build it.

Installing golang:

In rpm distributions:

You can quickly install golang on your os through the command line interface with:

sudo dnf -y install golang

Installing golang in deb distributions:

You can quickly install golang on your os through the command line interface with:

sudo apt-get install golang

Installing golang in Arch distribution:

You can quickly install golang on your os through the command line interface with:

sudo pacman -S golang 

2- Download the program from this link:

 https://notabug.org/alimiracle/irc-log-recorder
 
then extract it ( in home folder recomended ).

or if you have git  run this command from terminal :

git clone  https://notabug.org/alimiracle/irc-log-recorder

then go to program folder by typing:

  cd irc-log-recorder
  
then type:

chmod +x install

sudo ./install

to config the programme 

open /etc/ircconfig.txt as root
sudo nano /etc/ircconfig.txt
replace urukbot with your bot web page

replace http://myip.com/  with your server host

if your host is local
Type http://localhost/
replace 8080 with your server port You want your bot listen to it

replace myserver:444 with your irc server and port

replace myname with your bot name

replace myuser with your bot user name

replace mypassword  with your bot password

replace #emacs with your irc channel

then save the file

now you can run the bot 

type irl to run it

and You can reach the log from http://yourip:8080
replace yourip with your Computer IP

and if you have good Internet 
you can make all the world se the log by link your ip with domain 

Trouble: 

All efforts have been made to ensure the smooth and correct running of this application. 

If you find that irc log recorder is behaving abnormaly though, there are 3 options : 

1) Turn it off and run away. Not an option I would advise. 

2) Write a harsh comment that says how this app is a pile of **** and you can't believe I 
even dared to waste your time.. Again, not a great option, but it does make me laugh when I 
read some of the stuff.. :-) 

3) Send me a short email with error type and any other 
information you think is relevant, and I'll fix it. Jackpot. 

When I find a bug, I crush it. If I don't find it, and you do, and don't tell me, it lives and we all 
lose.. I'm not a mind-reader. Or a Computer-reader. I'm not one of the X-Men. 
i'm alimiracle
