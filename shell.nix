{ pkgs ? import <nixpkgs> { } }:
let
  my-python-packages = ps: with ps; [
    (
      buildPythonPackage rec {
        pname = "catppuccin";
        version = "1.2.0";
        src = fetchPypi {
          inherit pname version;
          sha256 = "sha256-hUNt6RHntQzamDX1SdhBzSj3pR/xxb6lpwzvYnqwOIo=";
        };
        doCheck = false;
        propagatedBuildInputs = [ ];
      }
    )
  ];
in
pkgs.mkShell {
  buildInputs = [
    (pkgs.python311.withPackages my-python-packages)
  ];
}