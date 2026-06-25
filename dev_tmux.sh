#!/bin/bash

SESSION="portfolio-dev"

tmux kill-session -t $SESSION 2>/dev/null

tmux new-session -d -s $SESSION -n dev

# Pane 1 (top-left): Go server (air)
tmux send-keys -t $SESSION:0.0 "air" C-m

# Split vertically → Pane 2 (top-right): templ
tmux split-window -h -t $SESSION:0.0
tmux send-keys -t $SESSION:0.1 "templ generate --watch" C-m

# Split bottom pane: Tailwind
tmux split-window -v -t $SESSION:0.1
tmux send-keys -t $SESSION:0.2 "tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch" C-m

# Focus main window
tmux select-layout -t $SESSION:0 tiled

# Attach session
tmux attach -t $SESSION
