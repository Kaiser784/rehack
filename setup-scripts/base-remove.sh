#!/bin/bash

#sublime text
sudo apt-get purge sublime-text
rm -rf ~/.config/sublime-text

#vscode
sudo apt purge code
sudo apt autoremove