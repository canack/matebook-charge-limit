#!/bin/sh

[ $(id -u) -eq 0 ]  && \
(echo "[!] Run without root permissions for now.") && \
(echo "[!] You will then be prompted for permission with sudo") && exit 1



GROUP_NAME=$(id -gn)
(sudo cp bin/chargelimit /usr/local/bin/chargelimit) && \
(echo "[+] Moved into /usr/local/bin for execute anywhere") || \
(echo "[!] An error occurred when moving /usr/local/bin" &&  exit 127)

(sudo chown root:"$GROUP_NAME" /usr/local/bin/chargelimit) && \
(echo "[+] Group permissions") || \
(echo "[!] An error occurred when setting group permissions" &&  exit 127)

(sudo chmod 4750 /usr/local/bin/chargelimit) && \
(echo "[+] Execute permissions" && echo "[+] Everything is done!") || \
(echo "[!] An error occurred when setting execute permissions" &&  exit 127)



