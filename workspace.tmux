#!/bin/bash

tmux-workspace "GoWas-ProjectTemplate" "main" -c "make run && nvim && zsh" \
	-w "builder" -c "make dev"
