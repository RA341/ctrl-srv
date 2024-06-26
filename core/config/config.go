package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	General      GeneralConfig
	Network      NetworkConfig
	Qbit         QbitConfig
	DiscordNotif DiscordNotifConfig
}

type GeneralConfig struct {
	AutoUpdate     bool
	UpdateInterval int
}

type NetworkConfig struct {
	Host string
	Port int
}

type QbitConfig struct {
	Enable bool
	Url    string
	User   string
	Pass   string
	SID    string
}

type DiscordNotifConfig struct {
	Enable     bool
	WebhookURL string
	Username   string
	AvatarURL  string
}

const defaultConfigFileName = "config.ini"

var config Config

func Get() *Config {
	return &config
}

func Load() {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to get current working directory: %w", err))
	}
	configPath := filepath.Join(cwd, defaultConfigFileName)

	CreateDefaultConfigIfNotExists(configPath)

	iniCfg, err := ini.Load(configPath)
	if err != nil {
		log.Fatal("Failed to load config file:", err)
	}

	err = parseINI(iniCfg)
	if err != nil {
		log.Fatal("Failed to parse config file:", err)
	}

}

func parseINI(cfg *ini.File) error {
	// General section
	generalSection := cfg.Section("General")
	config.General = GeneralConfig{
		AutoUpdate:     generalSection.Key("auto_update").MustBool(true),
		UpdateInterval: generalSection.Key("update_interval").MustInt(168),
	}

	// Network section
	networkSection := cfg.Section("Network")
	config.Network = NetworkConfig{
		Host: networkSection.Key("Host").String(),
		Port: networkSection.Key("Port").MustInt(9220),
	}
	validateNetworkSection()

	// Qbit section
	qbitSection := cfg.Section("Qbit")

	config.Qbit = QbitConfig{
		Enable: qbitSection.Key("enable").MustBool(true),
		Url:    createFullUrl(qbitSection.Key("host").MustString("NOHOST"), qbitSection.Key("port").MustInt(8085)),
		User:   qbitSection.Key("username").String(),
		Pass:   qbitSection.Key("password").String(),
		SID:    "",
	}
	validateQbitSection()

	// Discord notifications section
	discordSection := cfg.Section("notifications.Discord")
	config.DiscordNotif = DiscordNotifConfig{
		Enable:     discordSection.Key("enable").MustBool(true),
		WebhookURL: discordSection.Key("discord_webhook_url").String(),
		Username:   discordSection.Key("username").MustString("CTRL Bot"),
		AvatarURL:  discordSection.Key("avatar_url").String(),
	}
	validateDiscordSection()

	return nil
}

func createFullUrl(host string, port int) string {
	if strings.HasPrefix(host, "https") {
		// if host is https address do not add the port
		return host
	}
	return fmt.Sprintf("%s:%s", host, strconv.Itoa(port))
}

func CreateDefaultConfigIfNotExists(configPath string) {
	// Check if the file already exists
	if _, err := os.Stat(configPath); err == nil {
		// File exists, no need to create
		log.Print("Config file found")
		return
	} else if !os.IsNotExist(err) {
		// Some other error occurred
		log.Fatal(fmt.Errorf("error checking config file: %w", err))
	}

	// Create a new INI file
	cfg := ini.Empty()

	// [General] section
	secGeneral := createSection(cfg, "General")
	createKey(cfg, secGeneral, "auto_update", "true", "")
	createKey(cfg, secGeneral, "update_interval", "168", "how often to check for updates in hours (default is weekly)")

	// [Network] section
	secNetwork := createSection(cfg, "Network")
	secNetwork.Comment = "BE CAREFUL WHEN CHANGING THIS, IT MAY CAUSE THE SERVER TO BECOME INACCESSIBLE"
	createKey(cfg, secNetwork, "Host", "0.0.0.0", "Warning do not add 'http://' in front of the ips")
	createKey(cfg, secNetwork, "Port", "9220", "")

	// [Qbit] section
	secQbit := createSection(cfg, "Qbit")
	createKey(cfg, secQbit, "enable", "true", "")
	createKey(cfg, secQbit, "host", "http://127.0.0.1", "ip or hostname and port of your qbittorrent instance (remember to add https or http accordingly)")
	createKey(cfg, secQbit, "port", "8085", "")
	createKey(cfg, secQbit, "username", "", "")
	createKey(cfg, secQbit, "password", "", "Remember to surround the password with '\"' for eg \"password\" and doe not contain '#'")

	// [notifications.Discord] section
	secDiscord := createSection(cfg, "notifications.Discord")
	createKey(cfg, secDiscord, "enable", "true", "")
	createKey(cfg, secDiscord, "discord_webhook_url", "", "more info https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks")
	createKey(cfg, secDiscord, "username", "CTRL Bot", "")
	createKey(cfg, secDiscord, "avatar_url", "https://i.imgur.com/KEungv8.png", "")

	// Save the file
	if err := cfg.SaveTo(configPath); err != nil {
		log.Fatal(fmt.Errorf("failed to save default config file: %w", err))
	}

	fmt.Printf("Created default config file at: %s\nIf you see any errors below, fill the required fields first", configPath)
}

func createSection(cfg *ini.File, section string) *ini.Section {
	createdSection, err := cfg.NewSection(section)
	if err != nil {
		log.Fatal(fmt.Errorf("error creating section %s: %w", section, err))
	}
	return createdSection
}

func createKey(cfg *ini.File, section *ini.Section, key string, value string, comment string) {
	createdKey, err := section.NewKey(key, value)
	if err != nil {
		log.Fatal(fmt.Errorf("error creating key %s: %w", key, err))
	}
	createdKey.Comment = comment
}
