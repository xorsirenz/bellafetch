{
  inputs.flakelight-go.url = "github:chikof/flakelight-go";
  outputs = {flakelight-go, ...}:
    flakelight-go ./. {};
}
