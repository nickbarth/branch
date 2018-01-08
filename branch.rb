class Branch < Formula
  desc "branch is sample utility for switching branches using chooser."
  homepage "https://github.com/nickbarth/branch"
  url "https://github.com/nickbarth/branch/archive/v0.1-alpha.tar.gz"
  version "1.0.0"
  sha256 "955ba22959bfa181aa79810e4094713e95f679ff21059eff910977e265d33866"

  def install
    bin.install "branch"
  end
end
