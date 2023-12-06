# Deepgram Virtual Assistant

This is a demo application for implementing a digital assistant using the [Open Virtual Assistant](https://github.com/dvonthenen/open-virtual-assistant) project. This leverages Deepgram's [Go SDK](https://github.com/deepgram/deepgram-go-sdk) to provide the Speech-To-Text functionality to make pieces of this demo a reality.

## Running the Demo

The current list of assistant implementations:

- [Transcriber](https://github.com/deepgram-devs/virtual-assistant/tree/main/cmd/bin/dictation) - A slient assistant that acts as a scribe that will take notes of your dictation, then on command, will email you the current note
  - To run this demo, you need to configure the config.json file with SMTP settings for your internet provider:
  
  ```
  cd ./cmd/bin/dictation

  # setup the configuration
  cp config.json-ORG config.json
  vi config.json
  # fill in the settings below

  # set the EMAIL_SMTP_PASSWORD environment variable in your profile, then run:
  go run main.go

  # OR supply the environment variable on the command line
  # (this should only be used for evaluation purposes)
  # then run:
  EMAIL_SMTP_PASSWORD="YOUR_PASSWORD" go run main.go
  ```

## Development and Contributing

Interested in contributing? We ❤️ pull requests!

To make sure our community is safe for all, be sure to review and agree to our
[Code of Conduct](./CODE_OF_CONDUCT.md). Then see the
[Contribution](./CONTRIBUTING.md) guidelines for more information.

## Getting Help

We love to hear from you so if you have questions, comments or find a bug in the
project, let us know! You can either:

- [Open an issue](https://github.com/deepgram/[reponame]/issues/new) on this repository
- Ask a question, share the cool things you're working on, or see what else we have going on in our [Community Forum](https://github.com/orgs/deepgram/discussions/)
- Tweet at us! We're [@DeepgramAI on Twitter](https://twitter.com/DeepgramAI)

## Further Reading

Check out the Developer Documentation at [https://developers.deepgram.com/](https://developers.deepgram.com/)
