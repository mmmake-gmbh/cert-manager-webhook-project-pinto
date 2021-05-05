# Cert Manager Webhook for Pinto DNS

Cert Manager Webhook for Pinto DNS is a ACME webhook for [cert-manager](https://cert-manager.io/) allowing users to use [Pinto DNS](https://www.pinto.com/en/docs/pinto-dns/) for DNS01 challenge.

## Getting started

### Prerequisites

- A [Pinto Access Key and a Pinto Secret Key](https://www.pinto.com/en/docs/generate-api-keys/)
- A valid domain configured on [Pinto DNS](https://www.pinto.com/en/docs/pinto-dns/)
- A Kubernetes cluster (v1.19+ recommended)
- [Helm 3](https://helm.sh/) [installed](https://helm.sh/docs/intro/install/) on your computer
- [cert-manager](https://github.com/jetstack/cert-manager) deployed on the cluster:
```bash
kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.0.4/cert-manager.yaml
```

### Installing

Once everything is set up, you can now install the Pinto Webhook:
- Clone this repository: 
```bash
git clone https://gitlab.com/whizus/cert-manager-webhook-pinto.git
```

- Run:
```bash
helm install pinto-webhook deploy/pinto-webhook
```
- Alternatively, you can install the webhook with default credentials with: 
```bash
helm install pinto-webhook deploy/pinto-webhook --set secret.accessKey=<YOUR-ACCESS-KEY> --set secret.secretKey=<YOUR-SECRET_KEY>
```

The Pinto Webhook is now installed! :tada:

### How to use it

**Note**: It uses the [cert-manager webhook system](https://cert-manager.io/docs/configuration/acme/dns01/webhook/). Everything after the issuer is configured is just cert-manager. You can find out more in [their documentation](https://cert-manager.io/docs/usage/).

Now that the webhook is installed, here is how to use it.
Let's say you need a certificate for `example.com` (should be registered in Pinto DNS).

First step is to create a secret containing the Pinto Access and Secret keys. Create the `pinto-secret.yaml` file with the following content:
(Only needed if you don't have default credentials as seen above).
```yaml
apiVersion: v1
stringData:
  SCW_ACCESS_KEY: <YOUR-pinto-ACCESS-KEY>
  SCW_SECRET_KEY: <YOUR-pinto-SECRET-KEY>
kind: Secret
metadata:
  name: pinto-secret
type: Opaque
```

And run:
```bash
kubectl create -f pinto-secret.yaml
```

Next step is to create a cert-manager `Issuer`. Create a `issuer.yaml` file with the following content:
```yaml
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: my-pinto-issuer
spec:
  acme:
    email: my-user@example.com
    # this is the acme staging URL
    server: https://acme-staging-v02.api.letsencrypt.org/directory
    # for production use this URL instead
    # server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: my-pinto-private-key-secret
    solvers:
    - dns01:
        webhook:
          groupName: acme.pinto.com
          solverName: pinto
          config:
            # Only needed if you don't have default credentials as seen above.
            accessKeySecretRef:
              key: SCW_ACCESS_KEY
              name: pinto-secret
            secretKeySecretRef:
              key: SCW_SECRET_KEY
              name: pinto-secret
```

And run:
```bash
kubectl create -f issuer.yaml
```

Finally, you can now create the `Certificate` object for `example.com`. Create a `certificate.yaml` file with the following content:
```yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: example-com
spec:
  dnsNames:
  - example.com
  issuerRef:
    name: my-pinto-issuer
  secretName: example-com-tls
```

And run:
```bash
kubectl create -f certificate.yaml
```

After some seconds, you should see the certificate as ready:
```bash
$ kubectl get certificate example-com
NAME          READY   SECRET            AGE
example-com   True    example-com-tls   1m12s
```

Your certificate is now available in the `example-com-tls` secret!

## Integration testing

Before running the test, you need:
- A valid domain on Pinto DNS (here `example.com`)
- The variables `SCW_ACCESS_KEY` and `SCW_SECRET_KEY` valid and in the environment

In order to run the integration tests, run:
```bash
TEST_ZONE_NAME=example.com make test
```
