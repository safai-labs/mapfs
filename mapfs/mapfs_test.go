package mapfs_test

import (
	"code.cloudfoundry.org/lager/lagertest"

	"code.cloudfoundry.org/lager"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"code.cloudfoundry.org/mapfs/mapfs_fakes"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"code.cloudfoundry.org/mapfs/mapfs"
	"github.com/hanwen/go-fuse/fuse"
	"code.cloudfoundry.org/goshims/syscallshim/syscall_fake"
	"time"
)

var _ = Describe("mapfs", func() {
	var (
		logger lager.Logger
		mapFS pathfs.FileSystem
		uid, gid int64

		fakeFS *mapfs_fakes.FakeFileSystem
		fakeSyscall *syscall_fake.FakeSyscall
		context *fuse.Context
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("test-fs")
		fakeFS = &mapfs_fakes.FakeFileSystem{}
		fakeSyscall = &syscall_fake.FakeSyscall{}
		uid = 5
		gid = 10
		context = &fuse.Context{
			Owner: fuse.Owner{50, 100},
		}
	})

	JustBeforeEach(func() {
		mapFS = mapfs.NewMapFileSystem(uid, gid, fakeFS, fakeSyscall)
	})

	Context("when there is a mapfs", func() {
		BeforeEach(func() {
		})
		AfterEach(func() {
			Expect(fakeSyscall.SetregidCallCount()).To(Equal(1))
			Expect(fakeSyscall.SetreuidCallCount()).To(Equal(1))
		})

		Context(".Chmod", func() {
			It("passes the function through to the underlying filesystem unchanged", func() {
				mapFS.Chmod("foo", uint32(0777), context)

				Expect(fakeFS.ChmodCallCount()).To(Equal(1))
				name, mode, passedContext := fakeFS.ChmodArgsForCall(0)
				Expect(name).To(Equal("foo"))
				Expect(mode).To(Equal(uint32(0777)))
				Expect(passedContext).To(Equal(context))
			})
		})

		Context(".Chown", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.Chown("foo", uint32(50), uint32(100), context)

				Expect(fakeFS.ChownCallCount()).To(Equal(1))
			})
			It("maps to the spec'd uid and gid when chowning to the current user", func() {
				mapFS.Chown("foo", uint32(50), uint32(100), context)

				Expect(fakeFS.ChownCallCount()).To(Equal(1))
				_, uid, gid, _ := fakeFS.ChownArgsForCall(0)
				Expect(uid).To(Equal(uint32(5)))
				Expect(gid).To(Equal(uint32(10)))
			})
		})

		Context(".Utimens", func() {
			It("passes the function through to the underlying filesystem", func() {
				t := time.Now()
				mapFS.Utimens("foo", &t, &t, context)

				Expect(fakeFS.UtimensCallCount()).To(Equal(1))
			})
		})

		Context(".Truncate", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.Truncate("foo", uint64(50), context)

				Expect(fakeFS.TruncateCallCount()).To(Equal(1))
			})
		})

		Context(".Access", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.Access("foo", uint32(0777), context)

				Expect(fakeFS.AccessCallCount()).To(Equal(1))
			})
		})

		Context(".Link", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.Link("foo", "bar", context)

				Expect(fakeFS.LinkCallCount()).To(Equal(1))
			})
		})

		Context(".Mknod", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.Mknod("foo", uint32(0777), uint32(0777), context)

				Expect(fakeFS.MknodCallCount()).To(Equal(1))
			})
		})

		Context(".Mkdir", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.Mkdir("foo", uint32(0777), context)

				Expect(fakeFS.MkdirCallCount()).To(Equal(1))
			})
		})

		Context(".Rename", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.Rename("foo", "bar", context)

				Expect(fakeFS.RenameCallCount()).To(Equal(1))
			})
		})

		Context(".Rmdir", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.Rmdir("foo", context)

				Expect(fakeFS.RmdirCallCount()).To(Equal(1))
			})
		})

		Context(".Unlink", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.Unlink("foo", context)

				Expect(fakeFS.UnlinkCallCount()).To(Equal(1))
			})
		})

		Context(".GetXAttr", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.GetXAttr("foo", "bar", context)

				Expect(fakeFS.GetXAttrCallCount()).To(Equal(1))
			})
		})

		Context(".ListXAttr", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.ListXAttr("foo", context)

				Expect(fakeFS.ListXAttrCallCount()).To(Equal(1))
			})
		})

		Context(".RemoveXAttr", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.RemoveXAttr("foo", "bar", context)

				Expect(fakeFS.RemoveXAttrCallCount()).To(Equal(1))
			})
		})

		Context(".SetXAttr", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.SetXAttr("foo", "bar", []byte("baz"), 0, context)

				Expect(fakeFS.SetXAttrCallCount()).To(Equal(1))
			})
		})

		Context(".Open", func() {
			It("passes the function through to the underlying filesystem", func() {
				context := &fuse.Context{}
				mapFS.Open("foo", uint32(0777), context)

				Expect(fakeFS.OpenCallCount()).To(Equal(1))
			})
		})

		Context(".Create", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.Create("foo", 0, uint32(0777), context)

				Expect(fakeFS.CreateCallCount()).To(Equal(1))
			})
		})

		Context(".OpenDir", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.OpenDir("foo", context)

				Expect(fakeFS.OpenDirCallCount()).To(Equal(1))
			})
		})

		Context(".Symlink", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.Symlink("foo", "bar", context)

				Expect(fakeFS.SymlinkCallCount()).To(Equal(1))
			})
		})

		Context(".Readlink", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.Readlink("foo", context)

				Expect(fakeFS.ReadlinkCallCount()).To(Equal(1))
			})
		})

		Context(".StatFs", func() {
			It("passes the function through to the underlying filesystem", func() {
				mapFS.StatFs("foo")

				Expect(fakeFS.StatFsCallCount()).To(Equal(1))
			})
		})

		Context(".GetAttr", func() {
			It("maps the uid/gid back to the fuse context uid when it matches the mapped id", func() {
				attr := &fuse.Attr{}
				attr.Uid = uint32(uid)
				attr.Gid = uint32(gid)
				attr.Mode = uint32(0777)
				fakeFS.GetAttrReturns(attr, fuse.OK)
				ret, code := mapFS.GetAttr("foo", context)

				Expect(fakeFS.GetAttrCallCount()).To(Equal(1))
				Expect(code).To(Equal(fuse.OK))
				Expect(ret.Uid).To(Equal(context.Uid))
				Expect(ret.Gid).To(Equal(context.Gid))
				Expect(ret.Mode).To(Equal(uint32(0777)))
			})
		})

		Context(".OnMount", func() {
			It("passes through to the underlying fs", func() {
				mapFS.OnMount(nil)

				Expect(fakeFS.OnMountCallCount()).To(Equal(1))
			})
		})
	})
})
