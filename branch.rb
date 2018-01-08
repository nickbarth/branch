class Branch < Formula
  desc "branch is sample utility for switching branches using chooser."
  homepage "https://github.com/nickbarth/branch"
  url "https://github.com/nickbarth/branch/releases/download/v0.1-alpha/branch"
  version "0.0.1"
  sha256 "f7aa096445711e7bb7da113619e1598ded816e1531a494d7b610ce82d012d1d5"

  def install
    bin.install "branch"
  end
end
