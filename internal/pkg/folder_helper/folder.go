package folder_helper

import (
	"github.com/allanpk716/ChineseSubFinder/internal/pkg/get_access_time"
	"github.com/allanpk716/ChineseSubFinder/internal/pkg/global_value"
	"github.com/allanpk716/ChineseSubFinder/internal/pkg/log_helper"
	"github.com/allanpk716/ChineseSubFinder/internal/pkg/my_util"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// --------------------------------------------------------------
// Debug
// --------------------------------------------------------------

// GetRootDebugFolder 在程序的根目录新建，调试用文件夹
func GetRootDebugFolder() (string, error) {
	if global_value.DefDebugFolder == "" {
		nowProcessRoot, err := os.Getwd()
		if err != nil {
			return "", err
		}
		nowProcessRoot = filepath.Join(nowProcessRoot, cacheRootFolderName, DebugFolder)
		err = os.MkdirAll(nowProcessRoot, os.ModePerm)
		if err != nil {
			return "", err
		}
		global_value.DefDebugFolder = nowProcessRoot
		return nowProcessRoot, err
	}
	return global_value.DefDebugFolder, nil
}

// GetDebugFolderByName 根据传入的 strings (["aa", "bb"]) 组成  DebugFolder/aa/bb 这样的路径
func GetDebugFolderByName(names []string) (string, error) {

	rootPath, err := GetRootDebugFolder()
	if err != nil {
		return "", err
	}

	tmpFolderFullPath := rootPath
	for _, name := range names {
		tmpFolderFullPath = filepath.Join(tmpFolderFullPath, name)
	}
	err = os.MkdirAll(tmpFolderFullPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	return tmpFolderFullPath, nil
}

// CopyFiles2DebugFolder 把文件放入到 Debug 文件夹中，新建 desFolderName 文件夹
func CopyFiles2DebugFolder(names []string, subFiles []string) error {
	debugFolderByName, err := GetDebugFolderByName(names)
	if err != nil {
		return err
	}

	// 复制下载在 tmp 文件夹中的字幕文件到视频文件夹下面
	for _, subFile := range subFiles {
		newFn := filepath.Join(debugFolderByName, filepath.Base(subFile))
		err = my_util.CopyFile(subFile, newFn)
		if err != nil {
			return err
		}
	}

	return nil
}

// --------------------------------------------------------------
// Tmp
// --------------------------------------------------------------

// GetRootTmpFolder 在程序的根目录新建，取缓用文件夹，每一个视频的缓存将在其中额外新建子集文件夹
func GetRootTmpFolder() (string, error) {
	if global_value.DefTmpFolder == "" {
		nowProcessRoot, err := os.Getwd()
		if err != nil {
			return "", err
		}
		nowProcessRoot = filepath.Join(nowProcessRoot, cacheRootFolderName, TmpFolder)
		err = os.MkdirAll(nowProcessRoot, os.ModePerm)
		if err != nil {
			return "", err
		}
		global_value.DefTmpFolder = nowProcessRoot
		return nowProcessRoot, err
	}
	return global_value.DefTmpFolder, nil
}

// GetTmpFolderByName 获取缓存的文件夹，没有则新建
func GetTmpFolderByName(folderName string) (string, error) {
	rootPath, err := GetRootTmpFolder()
	if err != nil {
		return "", err
	}
	tmpFolderFullPath := filepath.Join(rootPath, folderName)
	err = os.MkdirAll(tmpFolderFullPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return tmpFolderFullPath, nil
}

// ClearTmpFolderByName 清理指定的缓存文件夹
func ClearTmpFolderByName(folderName string) error {

	nowTmpFolder, err := GetTmpFolderByName(folderName)
	if err != nil {
		return err
	}

	return ClearFolder(nowTmpFolder)
}

// ClearRootTmpFolder 清理缓存的根目录，将里面的子文件夹一并清理
func ClearRootTmpFolder() error {
	nowTmpFolder, err := GetRootTmpFolder()
	if err != nil {
		return err
	}

	pathSep := string(os.PathSeparator)
	files, err := os.ReadDir(nowTmpFolder)
	if err != nil {
		return err
	}
	for _, curFile := range files {
		fullPath := nowTmpFolder + pathSep + curFile.Name()
		if curFile.IsDir() {
			err = os.RemoveAll(fullPath)
			if err != nil {
				return err
			}
		} else {
			// 这里就是文件了
			err = os.Remove(fullPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// --------------------------------------------------------------
// Sub Fix Cache
// --------------------------------------------------------------

// GetRootSubFixCacheFolder 在程序的根目录新建，字幕时间校正的缓存文件夹
func GetRootSubFixCacheFolder() (string, error) {
	if global_value.DefSubFixCacheFolder == "" {
		nowProcessRoot, err := os.Getwd()
		if err != nil {
			return "", err
		}
		nowProcessRoot = filepath.Join(nowProcessRoot, cacheRootFolderName, SubFixCacheFolder)
		err = os.MkdirAll(nowProcessRoot, os.ModePerm)
		if err != nil {
			return "", err
		}
		global_value.DefSubFixCacheFolder = nowProcessRoot
		return nowProcessRoot, err
	}
	return global_value.DefSubFixCacheFolder, nil
}

// GetSubFixCacheFolderByName 获取缓存的文件夹，没有则新建
func GetSubFixCacheFolderByName(folderName string) (string, error) {
	rootPath, err := GetRootSubFixCacheFolder()
	if err != nil {
		return "", err
	}
	tmpFolderFullPath := filepath.Join(rootPath, folderName)
	err = os.MkdirAll(tmpFolderFullPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return tmpFolderFullPath, nil
}

// --------------------------------------------------------------
// Common
// --------------------------------------------------------------

// ClearFolder 清空文件夹
func ClearFolder(folderFullPath string) error {
	pathSep := string(os.PathSeparator)
	files, err := os.ReadDir(folderFullPath)
	if err != nil {
		return err
	}
	for _, curFile := range files {
		fullPath := folderFullPath + pathSep + curFile.Name()
		if curFile.IsDir() {
			err = os.RemoveAll(fullPath)
			if err != nil {
				return err
			}
		} else {
			// 这里就是文件了
			err = os.Remove(fullPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ClearFolderEx 清空文件夹，文件夹名称有特殊之处，Hour-min-Nanosecond 的命名方式
// 如果调用的时候，已存在的文件夹的时间 min < 5 那么则清理
func ClearFolderEx(folderFullPath string, overtime int) error {

	_, hour, minute, _ := my_util.GetNowTimeString()
	pathSep := string(os.PathSeparator)
	files, err := os.ReadDir(folderFullPath)
	if err != nil {
		return err
	}
	for _, curFile := range files {
		fullPath := folderFullPath + pathSep + curFile.Name()
		if curFile.IsDir() {

			parts := strings.Split(curFile.Name(), "-")
			if len(parts) == 3 {
				// 基本是符合了，倒是还是需要额外的判断是否时间超过了
				tmpHourStr := parts[0]
				tmpMinuteStr := parts[1]
				tmpHour, err := strconv.Atoi(tmpHourStr)
				if err != nil {
					// 如果不符合命名格式，直接删除
					err = os.RemoveAll(fullPath)
					if err != nil {
						return err
					}
					continue
				}
				tmpMinute, err := strconv.Atoi(tmpMinuteStr)
				if err != nil {
					// 如果不符合命名格式，直接删除
					err = os.RemoveAll(fullPath)
					if err != nil {
						return err
					}
					continue
				}
				// 判断时间
				if tmpHour != hour {
					// 如果不符合命名格式，直接删除
					err = os.RemoveAll(fullPath)
					if err != nil {
						return err
					}
					continue
				}
				// 超过 5 min
				if minute-overtime > tmpMinute {
					// 如果不符合命名格式，直接删除
					err = os.RemoveAll(fullPath)
					if err != nil {
						return err
					}
					continue
				}
			} else {
				// 如果不符合命名格式，直接删除
				err = os.RemoveAll(fullPath)
				if err != nil {
					return err
				}
			}
		} else {
			// 这里就是文件了
			err = os.Remove(fullPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// GetConfigRootDirFPath 获取 Config 的根目录，不同系统不一样
func GetConfigRootDirFPath() string {

	nowConfigFPath := ""
	sysType := runtime.GOOS
	if sysType == "linux" {
		nowConfigFPath = configDirRootFPathLinux
	} else if sysType == "windows" {
		nowConfigFPath = configDirRootFPathWindows
	} else if sysType == "darwin" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic("GetConfigRootDirFPath darwin get UserHomeDir, Error:" + err.Error())
		}
		nowConfigFPath = home + configDirRootFPathDarwin
	} else {
		panic("GetConfigRootDirFPath can't matched OSType: " + sysType + " ,You Should Implement It Yourself")
	}

	return nowConfigFPath
}

// ClearIdleSubFixCacheFolder 清理闲置的字幕修正缓存文件夹
func ClearIdleSubFixCacheFolder(rootSubFixCacheFolder string, outOfDate time.Duration) error {

	/*
		从 GetRootSubFixCacheFolder 目录下，遍历第一级目录中的文件夹
		然后每个文件夹中，统计里面最后的访问时间（可能有多个文件），如果超过某个时间范围就标记删除这个文件夹
	*/
	pathSep := string(os.PathSeparator)
	files, err := os.ReadDir(rootSubFixCacheFolder)
	if err != nil {
		return err
	}
	wait2ScanFolder := make([]string, 0)
	for _, curFile := range files {

		fullPath := rootSubFixCacheFolder + pathSep + curFile.Name()
		if curFile.IsDir() == true {
			// 需要关注文件夹
			wait2ScanFolder = append(wait2ScanFolder, fullPath)
		}
	}

	wait2DeleteFolder := make([]string, 0)
	getAccessTimeEx := get_access_time.GetAccessTimeEx{}
	cutOff := time.Now().Add(-outOfDate)
	for _, s := range wait2ScanFolder {

		files, err = os.ReadDir(s)
		if err != nil {
			return err
		}

		maxAccessTime := time.Now()
		// 需要统计这个文件夹下的所有文件的 AccessTIme，找出最新（最大的值）的那个时间，再比较
		for i, curFile := range files {

			fullPath := s + pathSep + curFile.Name()
			if curFile.IsDir() == true {
				continue
			}
			// 只需要关注文件
			accessTime, err := getAccessTimeEx.GetAccessTime(fullPath)
			if err != nil {
				return err
			}
			if i == 0 {
				maxAccessTime = accessTime
			}
			if my_util.Time2SecondNumber(accessTime) > my_util.Time2SecondNumber(maxAccessTime) {
				maxAccessTime = accessTime
			}
		}
		if maxAccessTime.Sub(cutOff) <= 0 {
			// 确认可以删除
			wait2DeleteFolder = append(wait2DeleteFolder, s)
		}
	}
	// 统一清理过期的文件夹
	for _, s := range wait2DeleteFolder {
		log_helper.GetLogger().Infoln("Try 2 clear SubFixCache Folder:", s)
		err := os.RemoveAll(s)
		if err != nil {
			return err
		}
	}

	return nil
}

// 缓存文件的位置信息，都是在程序的根目录下的 cache 中
const (
	cacheRootFolderName = "cache"           // 缓存文件夹总名称
	DebugFolder         = "CSF-DebugThings" // 调试相关的文件夹
	TmpFolder           = "CSF-TmpThings"   // 临时缓存的文件夹
	SubFixCacheFolder   = "CSF-SubFixCache" // 字幕时间校正的缓存文件夹，一般可以不清理
)

// 配置文件的位置信息，这个会根据系统版本做区分
const (
	configDirRootFPathWindows = "."                         // Windows 就是在当前的程序目录
	configDirRootFPathLinux   = "/config"                   // Linux 是在 /config 下
	configDirRootFPathDarwin  = "/.config/chinesesubfinder" // Darwin 是在 os.UserHomeDir()/.config/chinesesubfinder/ 下
)
