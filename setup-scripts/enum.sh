#!/bin/bash  

kdir = $(pwd)

#whois
sudo apt install whois -y

#sublist3r
sudo git clone https://github.com/aboul3la/Sublist3r.git /opt/Sublist3r
cd /opt/Sublist3r
pip install -r requirements.txt
#python dependencies
sudo pip install requests
sudo pip install dnspython
sudo pip install argparse
sudo ln -sfv /opt/Sublist3r/sublist3r.py /usr/bin/sublist3r
cd $kdir

#fping
sudo apt install fping -y

#nmap
sudo apt  install nmap -y

#dirb
sudo apt install dirb -y

#gobuster
sudo apt install gobuster -y
