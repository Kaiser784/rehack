#!/bin/bash  

#update
sudo apt-get update
sudo apt-get upgrade -y
sudo apt install snapd -y
sudo apt-get install git -y

#makes git ssh possible 
echo "Host github.com
User git
Hostname ssh.github.com
Port 443
" >> ~/.ssh/config

#sublime text
wget -qO - https://download.sublimetext.com/sublimehq-pub.gpg | sudo apt-key add -
sudo apt-get install apt-transport-https
echo "deb https://download.sublimetext.com/ apt/stable/" | sudo tee /etc/apt/sources.list.d/sublime-text.list
sudo apt-get update
sudo apt-get install sublime-text -y

#vscode
sudo apt install software-properties-common apt-transport-https wget
wget -q https://packages.microsoft.com/keys/microsoft.asc -O- | sudo apt-key add - 
sudo add-apt-repository "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main"
sudo apt install code -y

#pip3
sudo apt-get install python3-pip -y

