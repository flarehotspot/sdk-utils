default:
	node ./run.dev.js

build-docs:
	cd core/sdk/docs && zola build

serve-docs:
	cd core/sdk/docs && zola serve
