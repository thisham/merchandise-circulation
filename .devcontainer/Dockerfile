ARG variant=alpine3.14
FROM golang:1.17.5-${variant} AS builder

ARG username=vscode
ARG user_uid=1000
ARG user_gid=1000

# user
RUN adduser ${username} -s /bin/sh -D -u ${user_uid} ${user_gid} \
  && mkdir -p /etc/sudoers.d \
  && echo ${username} ALL=(root) NOPASSWD:ALL > /etc/sudoers.d/${username} \
  && chmod 0440 /etc/sudoers.d/${username}

# Setup shell
USER ${username}
RUN sh -c “$(wget -O- https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)” “” --unattended &> /dev/null

ENV ENV=“/home/$USERNAME/.ashrc” \
  ZSH=/home/$USERNAME/.oh-my-zsh \
  EDITOR=vi \
  LANG=en_US.UTF-8

RUN printf 'ZSH_THEME="robbyrussell"\nENABLE_CORRECTION="false"\nplugins=(git copyfile extract colorize dotenv encode64 golang)\nsource $ZSH/oh-my-zsh.sh' > "/home/$USERNAME/.zshrc"

RUN echo "exec which zsh" > "/home/$USERNAME/.zshrc"

USER root
