package facade

// 外观模式

// 外观角色

// HomeTheaterFacade 是外观类，提供简化的接口
type HomeTheaterFacade struct {
	dvdPlayer   *DVDPlayer
	projector   *Projector
	soundSystem *SoundSystem
}

// NewHomeTheaterFacade 创建外观实例
func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		dvdPlayer:   &DVDPlayer{},
		projector:   &Projector{},
		soundSystem: &SoundSystem{},
	}
}

// WatchMovie 封装启动家庭影院的操作
func (h *HomeTheaterFacade) WatchMovie(movie string) string {
	result := "开始观看电影...\n"
	h.dvdPlayer.TurnOn()
	h.projector.TurnOn()
	h.projector.SetInput("DVD")
	h.soundSystem.TurnOn()
	h.soundSystem.SetVolume(10)
	h.dvdPlayer.Play(movie)
	return result
}

// EndMovie 封装关闭家庭影院的操作
func (h *HomeTheaterFacade) EndMovie() string {
	result := "结束电影播放...\n"
	h.dvdPlayer.TurnOff()
	h.projector.TurnOff()
	h.soundSystem.TurnOff()
	return result
}
