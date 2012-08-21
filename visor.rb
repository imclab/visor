require 'formula'

class Visor < Formula
  homepage 'http://github.com/soundcloud/visor'
  url 'https://github.com/soundcloud/visor/zipball/master'
  depends_on 'go'
  skip_clean 'bin'
  version '0.5.4'


  def install
    begin
      system("which hg")
    rescue
      system "brew install hg"
    end
    system "make GOPATH=#{buildpath} GOBIN=#{prefix}/bin"
  end

  def test
    system "visor --version"
  end
end
