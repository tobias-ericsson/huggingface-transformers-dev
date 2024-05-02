# L&L (Lunch & Learn)

#### Useful tools bringing Akind of A.I. from the clouds, into your hands while having lunch

## What are we doing?

[Silicon Valley hotdog or not](https://www.youtube.com/watch?v=vIci3C4JkL0&t=37s)

## How are we doing it?

### AWESOME

Akind Platform and the Akind Wonderful Enterprise Service Orchestration and Management Engine [AWESOME].

https://docs.akind.tech/docs/how-to/onboarding/

https://docs.akind.tech/docs/how-to/install-cli/

https://docs.akind.tech/docs/tutorials/deploy-your-first-service/

```
aw login
aw init hotdog-or-not
```
==> https://github.com/CrowdCollective/hotdog-or-not

### Codespaces

[.devcontainer/devcontainer.json](.devcontainer/devcontainer.json)

To rebuild container `Shift+Command+P >rebuild`

### Taskfile

https://taskfile.dev/

[Taskfile.yml](Taskfile.yml)

### Huggingface transformers.js

Huggingface - a library, platform and community for machine learning.

State-of-the-art Machine Learning for the web. Run 🤗 Transformers directly in your browser, with no need for a server!

Transformers.js uses [ONNX Runtime](https://onnxruntime.ai/) to run models in the browser.

ONNX Runtime is an open-source software framework designed to accelerate machine learning (ML) model inference, it's 
a cross-platform runtime environment that can execute models created in various deep learning frameworks.

Machine learning inference is the process of using a trained model to make prediction on new data.

https://huggingface.co/docs/transformers.js/en/index


### More AWESOME

```
aw status
aw access:enable -s hotdog-or-not -e dev
aw log -s hotdog-or-not -e dev -t
```

Sensible defaults:

* Dockerfile
* Listen on port :80
* Health check on /health

https://hotdog-or-not.dev.akind.tech/

