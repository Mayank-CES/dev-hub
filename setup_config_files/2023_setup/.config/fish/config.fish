if status is-interactive
    # Commands to run in interactive sessions can go here
end
# Initialize rbenv
set -x PATH $HOME/.rbenv/bin $PATH
status --is-interactive; and rbenv init - | source

set -x GITHUB_USERNAME "Mayank-CES"
set -x GITHUB_ACCESS_TOKEN ""
set -gx LDFLAGS "-L/opt/homebrew/opt/libpq/lib"
set -gx CPPFLAGS "-I/opt/homebrew/opt/libpq/include"
set -gx PKG_CONFIG_PATH "/opt/homebrew/opt/libpq/lib/pkgconfig"
