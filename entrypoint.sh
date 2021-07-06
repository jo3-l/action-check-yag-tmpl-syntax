cp /check_yag_tmpl_syntax.json "$HOME/"
echo "::add-matcher::$HOME/check_yag_tmpl_syntax.json"
./app
exit $?