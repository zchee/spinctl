// Copyright 2019 The spinctl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const defaultBoilerPlate = `
# Copyright 2019 The spinctl Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
`

var (
	completionLong = `
		Output shell completion code for the specified shell (bash or zsh).
		The shell code must be evaluated to provide interactive
		completion of spinctl commands.  This can be done by sourcing it from
		the .bash_profile.

		Note for zsh users: [1] zsh completions are only supported in versions of zsh >= 5.2`

	completionExample = `
		# Installing bash completion on macOS using homebrew
		## If running Bash 3.2 included with macOS
		    brew install bash-completion
		## or, if running Bash 4.1+
		    brew install bash-completion@2
		## If spinctl is installed via homebrew, this should start working immediately.
		## If you've installed via other means, you may need add the completion to your completion directory
		    spinctl completion bash > $(brew --prefix)/etc/bash_completion.d/spinctl


		# Installing bash completion on Linux
		## If bash-completion is not installed on Linux, please install the 'bash-completion' package
		## via your distribution's package manager.
		## Load the spinctl completion code for bash into the current shell
		    source <(spinctl completion bash)
		## Write bash completion code to a file and source if from .bash_profile
		    spinctl completion bash > /etc/bash_completion.d/spinctl.bash.inc
		    printf "
		      # spinctl shell completion
		      source '/etc/bash_completion.d/spinctl.bash.inc'
		      " >> $HOME/.bash_profile
		    source $HOME/.bash_profile

		# Load the spinctl completion code for zsh[1] into the current shell
		    source <(spinctl completion zsh)
		# Set the spinctl completion code for zsh[1] to autoload on startup
		    spinctl completion zsh > "${fpath[1]}/_spinctl"`
)

var (
	completionShells = map[string]func(out io.Writer, cmd *cobra.Command) error{
		"bash": runCompletionBash,
		"zsh":  runCompletionZsh,
	}
)

// NewCmdCompletion creates the spinctl shell completion.
func NewCmdCompletion(out io.Writer) *cobra.Command {
	shells := []string{}
	for s := range completionShells {
		shells = append(shells, s)
	}

	cmd := &cobra.Command{
		Use:                   "completion SHELL",
		DisableFlagsInUseLine: true,
		Short:                 "Output shell completion code for the specified shell (bash or zsh)",
		Long:                  completionLong,
		Example:               completionExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCompletion(out, cmd, args)
		},
		ValidArgs: shells,
	}

	return cmd
}

func runCompletion(out io.Writer, cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("shell not specified")
	}
	if len(args) > 1 {
		return errors.New("too many arguments. expected only the shell type")
	}
	runFunc, found := completionShells[args[0]]
	if !found {
		return errors.Errorf("Unsupported shell type %q", args[0])
	}

	return runFunc(out, cmd.Parent())
}

func runCompletionBash(out io.Writer, cmd *cobra.Command) error {
	if _, err := out.Write([]byte(defaultBoilerPlate)); err != nil {
		return err
	}

	return cmd.GenBashCompletion(out)
}

func runCompletionZsh(out io.Writer, cmd *cobra.Command) error {
	zshHead := "#compdef spinctl\n"

	out.Write([]byte(zshHead))

	if _, err := out.Write([]byte(defaultBoilerPlate)); err != nil {
		return err
	}

	zshInitialization := `
__spinctl_bash_source() {
	alias shopt=':'
	alias _expand=_bash_expand
	alias _complete=_bash_comp
	emulate -L sh
	setopt kshglob noshglob braceexpand

	source "$@"
}

__spinctl_type() {
	# -t is not supported by zsh
	if [ "$1" == "-t" ]; then
		shift

		# fake Bash 4 to disable "complete -o nospace". Instead
		# "compopt +-o nospace" is used in the code to toggle trailing
		# spaces. We don't support that, but leave trailing spaces on
		# all the time
		if [ "$1" = "__spinctl_compopt" ]; then
			echo builtin
			return 0
		fi
	fi
	type "$@"
}

__spinctl_compgen() {
	local completions w
	completions=( $(compgen "$@") ) || return $?

	# filter by given word as prefix
	while [[ "$1" = -* && "$1" != -- ]]; do
		shift
		shift
	done
	if [[ "$1" == -- ]]; then
		shift
	fi
	for w in "${completions[@]}"; do
		if [[ "${w}" = "$1"* ]]; then
			echo "${w}"
		fi
	done
}

__spinctl_compopt() {
	true # don't do anything. Not supported by bashcompinit in zsh
}

__spinctl_ltrim_colon_completions()
{
	if [[ "$1" == *:* && "$COMP_WORDBREAKS" == *:* ]]; then
		# Remove colon-word prefix from COMPREPLY items
		local colon_word=${1%${1##*:}}
		local i=${#COMPREPLY[*]}
		while [[ $((--i)) -ge 0 ]]; do
			COMPREPLY[$i]=${COMPREPLY[$i]#"$colon_word"}
		done
	fi
}

__spinctl_get_comp_words_by_ref() {
	cur="${COMP_WORDS[COMP_CWORD]}"
	prev="${COMP_WORDS[${COMP_CWORD}-1]}"
	words=("${COMP_WORDS[@]}")
	cword=("${COMP_CWORD[@]}")
}

__spinctl_filedir() {
	local RET OLD_IFS w qw

	__spinctl_debug "_filedir $@ cur=$cur"
	if [[ "$1" = \~* ]]; then
		# somehow does not work. Maybe, zsh does not call this at all
		eval echo "$1"
		return 0
	fi

	OLD_IFS="$IFS"
	IFS=$'\n'
	if [ "$1" = "-d" ]; then
		shift
		RET=( $(compgen -d) )
	else
		RET=( $(compgen -f) )
	fi
	IFS="$OLD_IFS"

	IFS="," __spinctl_debug "RET=${RET[@]} len=${#RET[@]}"

	for w in ${RET[@]}; do
		if [[ ! "${w}" = "${cur}"* ]]; then
			continue
		fi
		if eval "[[ \"\${w}\" = *.$1 || -d \"\${w}\" ]]"; then
			qw="$(__spinctl_quote "${w}")"
			if [ -d "${w}" ]; then
				COMPREPLY+=("${qw}/")
			else
				COMPREPLY+=("${qw}")
			fi
		fi
	done
}

__spinctl_quote() {
    if [[ $1 == \'* || $1 == \"* ]]; then
        # Leave out first character
        printf %q "${1:1}"
    else
	printf %q "$1"
    fi
}

autoload -U +X bashcompinit && bashcompinit

# use word boundary patterns for BSD or GNU sed
LWORD='[[:<:]]'
RWORD='[[:>:]]'
if sed --help 2>&1 | grep -q GNU; then
	LWORD='\<'
	RWORD='\>'
fi

__spinctl_convert_bash_to_zsh() {
	sed \
	-e 's/declare -F/whence -w/' \
	-e 's/_get_comp_words_by_ref "\$@"/_get_comp_words_by_ref "\$*"/' \
	-e 's/local \([a-zA-Z0-9_]*\)=/local \1; \1=/' \
	-e 's/flags+=("\(--.*\)=")/flags+=("\1"); two_word_flags+=("\1")/' \
	-e 's/must_have_one_flag+=("\(--.*\)=")/must_have_one_flag+=("\1")/' \
	-e "s/${LWORD}_filedir${RWORD}/__spinctl_filedir/g" \
	-e "s/${LWORD}_get_comp_words_by_ref${RWORD}/__spinctl_get_comp_words_by_ref/g" \
	-e "s/${LWORD}__ltrim_colon_completions${RWORD}/__spinctl_ltrim_colon_completions/g" \
	-e "s/${LWORD}compgen${RWORD}/__spinctl_compgen/g" \
	-e "s/${LWORD}compopt${RWORD}/__spinctl_compopt/g" \
	-e "s/${LWORD}declare${RWORD}/builtin declare/g" \
	-e "s/\\\$(type${RWORD}/\$(__spinctl_type/g" \
	<<'BASH_COMPLETION_EOF'
`
	out.Write([]byte(zshInitialization))

	buf := new(bytes.Buffer)
	cmd.GenBashCompletion(buf)
	out.Write(buf.Bytes())

	zshTail := `
BASH_COMPLETION_EOF
}

__spinctl_bash_source <(__spinctl_convert_bash_to_zsh)
_complete spinctl 2>/dev/null
`
	out.Write([]byte(zshTail))
	return nil
}
