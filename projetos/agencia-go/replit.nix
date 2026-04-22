{ pkgs }: pkgs.mkShell {
  packages = with pkgs; [
    go_1_21
  ];
}