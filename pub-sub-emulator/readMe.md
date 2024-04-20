# Install PubSub emulator

```
docker run --rm -p 8085:8085 google/cloud-sdk:emulators /bin/bash -c "gcloud beta emulators pubsub start --project=core-financing-id --host-port='0.0.0.0:8085'"
```

## Add to env on your apps

```
export PUBSUB_EMULATOR_HOST=localhost:8085
```

# Clone this apps
```
git clone https://github.com/NeoScript/pubsub-ui.git
```
# Run UI apps
```
cd pubsub-ui
code .
```

# UI 
![Screenshot 2024-04-20 at 10.41.38.png](..%2F..%2F..%2F..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ff3%2Fhl0dh0h51wg5jfq9kh0h7d8c0000gn%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_HD8Hys%2FScreenshot%202024-04-20%20at%2010.41.38.png)