include $(GOROOT)/src/Make.inc

TARG=slide
CLEANFILES+=$(TARG).exe _test _testmain.go test.out build.out

all:$(TARG)
$(TARG):main.$O
	$(LD) -Lgame/_obj -o $@ $^
%.$O:%.go
	$(GC) -Igame/_obj -o $@ $^
game:
	$(MAKE) -C game
.PHONY:game

