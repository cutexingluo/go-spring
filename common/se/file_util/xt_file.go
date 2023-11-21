package file_util

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
)

// GetProjectAbsPath returns the project absolute path
func GetProjectAbsPath() (path string, ok error) {
	return filepath.Abs("")
}

// GetCurrentFileAbsPath returns the current file absolute path
func GetCurrentFileAbsPath() (string, bool) {
	_, file, _, ok := runtime.Caller(1) // return caller location
	return file, ok
}

type File struct {
	Path string // package path
}

func NewFile(path string) *File {
	return &File{Path: path}
}

// GetAbsPath returns the  absolute path
func (_this *File) GetAbsPath() (string, error) {
	return filepath.Abs(_this.Path)
}

// IsAbs returns whether the path is an absolute path
func (_this *File) IsAbs() bool {
	return filepath.IsAbs(_this.Path)
}

// GetDirPath returns the dir  absolute path
func (_this *File) GetDirPath() (string, error) {
	return filepath.Abs(_this.Path)
}

// IsFile returns check whether the file is a file
func (_this *File) IsFile() (bool, error) {
	dir, err := _this.IsDir()
	if err != nil {
		return false, err
	}
	return !dir, nil
}

// IsDir returns  check whether the file is a directory
func (_this *File) IsDir() (bool, error) {
	fileInfo, err := os.Stat(_this.Path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

// GetFileName returns the  absolute path
func (_this *File) GetFileName() string {
	return filepath.Base(_this.Path)
}

// GetFileExt  returns the  absolute path
func (_this *File) GetFileExt() string {
	return filepath.Ext(_this.Path)
}

// Open returns the os.File
func (_this *File) Open() (*os.File, error) {
	return os.Open(_this.Path)
}

// OpenFunc use the os.Open function to open the file
func (_this *File) OpenFunc(action func(*os.File) error) (err error) {
	file, err := os.Open(_this.Path)
	defer func(file *os.File) {
		err = file.Close()
	}(file)
	if err != nil {
		return err
	}
	err = action(file)
	return err
}

// OpenFile returns the os.OpenFile
func (_this *File) OpenFile(flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(_this.Path, flag, perm)
}

// OpenFileFunc use the os.OpenFile function to open the file
func (_this *File) OpenFileFunc(flag int, perm os.FileMode, action func(*os.File) error) (err error) {
	file, err := os.OpenFile(_this.Path, flag, perm)
	defer func(file *os.File) {
		err = file.Close()
	}(file)
	if err != nil {
		return err
	}
	err = action(file)
	return err
}

// WalkDir walk the fileDirPath
func (_this *File) WalkDir(f func(path string, info os.FileInfo, err error) error) (err error) {
	dirPath, err := _this.GetDirPath()
	if err != nil {
		return err
	}
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		return f(path, info, err)
	})
	return err
}

// ParseFile parse the file
func (_this *File) ParseFile() (f *ast.File, err error) {
	fSet := token.NewFileSet()
	return parser.ParseFile(fSet, _this.Path, nil, parser.AllErrors)
}

// Inspect the file
func (_this *File) Inspect(f func(n ast.Node) bool) (err error) {
	file, err := _this.ParseFile()
	if err != nil {
		return err
	}
	ast.Inspect(file, f)
	return nil
}

// InspectMethods inspect methods
func (_this *File) InspectMethods(f func(funcDecl *ast.FuncDecl) bool) (err error) {
	return _this.Inspect(func(n ast.Node) bool {
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			return f(funcDecl)
		}
		return true
	})
}

// GetDirAllFiles  returns all the files in the fileDirPath
func (_this *File) GetDirAllFiles() ([]string, error) {
	dirPath, err := _this.GetDirPath()
	if err != nil {
		return nil, err
	}
	return GetAllFiles(dirPath)
}

//---------------public methods ------------------

// GetAllFiles 获取指定目录下的所有文件
func GetAllFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files, err
}

// ParseFile 解析文件
func ParseFile(filePath string, f func(n ast.Node) bool) (err error) {
	fSet := token.NewFileSet()
	file, err := parser.ParseFile(fSet, filePath, nil, parser.AllErrors)
	if err != nil {
		//fmt.Println("Error parsing file:", err)
		return err
	}
	ast.Inspect(file, func(n ast.Node) bool {
		return f(n)
	})
	return err
}

// BuildFunctionCallExpr 构建函数调用表达式
func BuildFunctionCallExpr(pkgName string, funcName string, args []ast.Expr) *ast.CallExpr {
	selectorExpr := &ast.SelectorExpr{ // 构建函数选择器
		X:   &ast.Ident{Name: pkgName},  // 替换为实际的包名或包导入路径
		Sel: &ast.Ident{Name: funcName}, // 构建函数标识符
	}
	// 构建函数调用表达式
	callExpr := &ast.CallExpr{
		Fun:  selectorExpr,
		Args: args,
	}
	return callExpr
}
