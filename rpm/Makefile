TGT=dial

RPM_VERSION = 1.0.o
RPM_RELEASE = master
RPM_TOP_DIR = $(shell rpm -E %{_topdir})
PRJHOME = $(shell pwd)

.PHONY: rpmclean 

$(TGT):

rpm:
	@echo [RPM] ; \
    	sed -e "s/@VERSION@/$(RPM_VERSION)/g" -e "s/@RELEASE@/$(RPM_RELEASE)/g" $(TGT).spec > ${RPM_TOP_DIR}/SPECS/$(TGT).spec ; \
    	cp -a -r ${PRJHOME} /tmp/$(TGT)-$(RPM_VERSION) ; \
    	cd /tmp ; \
    	tar zcvf $(RPM_TOP_DIR)/SOURCES/$(TGT)-$(RPM_VERSION).tar.gz $(TGT)-$(RPM_VERSION) ; \
    	rm -rf $(TGT)-$(RPM_VERSION) ; \
    	rpmbuild -bb $(RPM_TOP_DIR)/SPECS/$(TGT).spec ; \

rpmclean:	
	cp -r ~/rpmbuild/RPMS/x86_64/$(TGT)*$(RPM_VERSION)* ./  
	rm -rf ~/rpmbuild/SOURCES/$(TGT)* \
	~/rpmbuild/BUILD/$(TGT)* \
	~/rpmbuild/RPMS/x86_64/$(TGT)* \
	~/rpmbuild/SPEC/$(TGT)* 



