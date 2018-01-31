class Branch < Formula
  desc "branch is sample utility for switching branches using chooser."
  homepage "https://github.com/nickbarth/branch"
  url "https://github.com/nickbarth/branch/releases/download/v0.2-alpha/branch"
  version "0.2.0"
  sha256 "8acc76cc03fb4d991bbeabdda39ae51d8b6bd12cfcd3b7dbc41b74af1d3df9a7"

  def install
    bin.install "branch"
  end
end
