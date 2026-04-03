{
  inputs.flakelight-go.url = "github:chikof/flakelight-go";
  outputs = {
    self,
    flakelight-go,
    ...
  }:
    flakelight-go ./. {
      package = pkgs: let
        version = self.rev or self.dirtyRev or "dev";
        pciIds = "${pkgs.hwdata}/share/hwdata/pci.ids";
        amdgpuIds = "${pkgs.libdrm}/share/libdrm/amdgpu.ids";
        extraLdflags = "-X github.com/xorsirenz/bellafetch/pkg/linux.CUSTOM_PCI_IDS_PATH=${pciIds} -X github.com/xorsirenz/bellafetch/pkg/linux.CUSTOM_AMDGPU_IDS_PATH=${amdgpuIds}";
      in
        pkgs.stdenv.mkDerivation {
          pname = "bellafetch";
          inherit version;
          src = ./.;

          nativeBuildInputs = [
            pkgs.go
            pkgs.gnumake
          ];

          buildPhase = ''
            runHook preBuild

            export HOME="$TMPDIR"
            export XDG_CACHE_HOME="$TMPDIR/.cache"
            export GOCACHE="$TMPDIR/go-build"
            export GOMODCACHE="$TMPDIR/go-mod"

            make build \
              VERSION=${version} \
              EXTRA_LDFLAGS=${pkgs.lib.escapeShellArg extraLdflags}

            runHook postBuild
          '';

          installPhase = ''
            runHook preInstall
            mkdir -p "$out/bin"
            install -m755 bellafetch "$out/bin/bellafetch"
            runHook postInstall
          '';
        };
    };
}
