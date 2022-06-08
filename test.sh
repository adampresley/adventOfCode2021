	if grep -i 'API_KEY\s*=\s*["`]*[^"`]' README.md; then
		echo "It looks like you are attempting to set an API key in $file. This is not allowed."
		exit 1
	fi

