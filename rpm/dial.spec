Name: dial		
Version: @VERSION@
Release: @RELEASE@%{dist}
Summary: YAMU dial

Group: Applications/Internet
License: GPL
URL: http://www.yamutech.com
Source0: dial-%{version}.tar.gz
BuildRoot: %{_tmppath}/%{name}-%{version}-%{release}

%description
YAMU dial


%prep
%setup -q -n dial-%{version}


%build
#make %{?_smp_mflags}
make

%install
rm -rf %{buildroot}
mkdir -p %{buildroot}/usr/local/dial/
mkdir -p %{buildroot}/etc/
mkdir -p %{buildroot}/etc/init.d/

install -m 755 dial %{buildroot}/usr/local/dial/
install -m 755 config/diald  %{buildroot}/etc/init.d/

%clean
rm -rf %{buildroot}

%post
chkconfig --add diald

%files
%defattr(-,root,root,-)
/usr/local/dial/dial
/etc/init.d/diald

%preun
chkconfig --del diald

%changelog

%define __debug_install_post   \
	%{_rpmconfigdir}/find-debuginfo.sh %{?_find_debuginfo_opts} "%{_builddir}/%{?buildsubdir}"\
%{nil}
