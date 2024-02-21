package gopom

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
)

func Parse(path string) (*Project, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, _ := ioutil.ReadAll(file)
	var project Project

	err = xml.Unmarshal(b, &project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (p *Project) Marshal() ([]byte, error) {
	return xml.MarshalIndent(ProjectMarshal(*p), "", "    ")
}

type Project struct {
	XMLName                xml.Name                `xml:"project,omitempty"`
	Xmlns                  string                  `xml:"xmlns,attr,omitempty"`
	Xsi                    string                  `xml:"xsi,attr,omitempty"`
	SchemaLocation         string                  `xml:"schemaLocation,attr,omitempty"`
	ModelVersion           string                  `xml:"modelVersion,omitempty"`
	GroupID                string                  `xml:"groupId,omitempty"`
	ArtifactID             string                  `xml:"artifactId,omitempty"`
	Version                string                  `xml:"version,omitempty"`
	Packaging              string                  `xml:"packaging,omitempty"`
	Name                   string                  `xml:"name,omitempty"`
	Description            string                  `xml:"description,omitempty"`
	URL                    string                  `xml:"url,omitempty"`
	InceptionYear          string                  `xml:"inceptionYear,omitempty"`
	Organization           *Organization           `xml:"organization,omitempty"`
	Licenses               *[]License              `xml:"licenses>license,omitempty"`
	Developers             *[]Developer            `xml:"developers>developer,omitempty"`
	Contributors           *[]Contributor          `xml:"contributors>contributor,omitempty"`
	MailingLists           *[]MailingList          `xml:"mailingLists>mailingList,omitempty"`
	Prerequisites          *Prerequisites          `xml:"prerequisites,omitempty"`
	Properties             *Properties             `xml:"properties,omitempty"`
	Parent                 *Parent                 `xml:"parent,omitempty"`
	Modules                *[]string               `xml:"modules>module,omitempty"`
	SCM                    *Scm                    `xml:"scm,omitempty"`
	IssueManagement        *IssueManagement        `xml:"issueManagement,omitempty"`
	CIManagement           *CIManagement           `xml:"ciManagement,omitempty"`
	DistributionManagement *DistributionManagement `xml:"distributionManagement,omitempty"`
	DependencyManagement   *DependencyManagement   `xml:"dependencyManagement,omitempty"`
	Dependencies           *[]Dependency           `xml:"dependencies>dependency,omitempty"`
	Repositories           *[]Repository           `xml:"repositories>repository,omitempty"`
	PluginRepositories     *[]PluginRepository     `xml:"pluginRepositories>pluginRepository,omitempty"`
	Build                  *Build                  `xml:"build,omitempty"`
	Reporting              *Reporting              `xml:"reporting,omitempty"`
	Profiles               *[]Profile              `xml:"profiles>profile,omitempty"`
}

type ProjectMarshal struct {
	XMLName                xml.Name                `xml:"project,omitempty"`
	Xmlns                  string                  `xml:"xmlns,attr,omitempty"`
	Xsi                    string                  `xml:"xmlns:xsi,attr,omitempty"`
	SchemaLocation         string                  `xml:"xsi:schemaLocation,attr,omitempty"`
	ModelVersion           string                  `xml:"modelVersion,omitempty"`
	GroupID                string                  `xml:"groupId,omitempty"`
	ArtifactID             string                  `xml:"artifactId,omitempty"`
	Version                string                  `xml:"version,omitempty"`
	Packaging              string                  `xml:"packaging,omitempty"`
	Name                   string                  `xml:"name,omitempty"`
	Description            string                  `xml:"description,omitempty"`
	URL                    string                  `xml:"url,omitempty"`
	InceptionYear          string                  `xml:"inceptionYear,omitempty"`
	Organization           *Organization           `xml:"organization,omitempty"`
	Licenses               *[]License              `xml:"licenses>license,omitempty"`
	Developers             *[]Developer            `xml:"developers>developer,omitempty"`
	Contributors           *[]Contributor          `xml:"contributors>contributor,omitempty"`
	MailingLists           *[]MailingList          `xml:"mailingLists>mailingList,omitempty"`
	Prerequisites          *Prerequisites          `xml:"prerequisites,omitempty"`
	Properties             *Properties             `xml:"properties,omitempty"`
	Parent                 *Parent                 `xml:"parent,omitempty"`
	Modules                *[]string               `xml:"modules>module,omitempty"`
	SCM                    *Scm                    `xml:"scm,omitempty"`
	IssueManagement        *IssueManagement        `xml:"issueManagement,omitempty"`
	CIManagement           *CIManagement           `xml:"ciManagement,omitempty"`
	DistributionManagement *DistributionManagement `xml:"distributionManagement,omitempty"`
	DependencyManagement   *DependencyManagement   `xml:"dependencyManagement,omitempty"`
	Dependencies           *[]Dependency           `xml:"dependencies>dependency,omitempty"`
	Repositories           *[]Repository           `xml:"repositories>repository,omitempty"`
	PluginRepositories     *[]PluginRepository     `xml:"pluginRepositories>pluginRepository,omitempty"`
	Build                  *Build                  `xml:"build,omitempty"`
	Reporting              *Reporting              `xml:"reporting,omitempty"`
	Profiles               *[]Profile              `xml:"profiles>profile,omitempty"`
}

type Properties struct {
	Entries map[string]string
}

func (p *Properties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	type entry struct {
		XMLName xml.Name
		Key     string `xml:"name,attr"`
		Value   string `xml:",chardata"`
	}
	e := entry{}
	p.Entries = map[string]string{}
	for err = d.Decode(&e); err == nil; err = d.Decode(&e) {
		e.Key = e.XMLName.Local
		p.Entries[e.Key] = e.Value
	}
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (p Properties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tokens := []xml.Token{start}

	for k, v := range p.Entries {
		t := xml.StartElement{Name: xml.Name{"", k}}
		tokens = append(tokens, t, xml.CharData(v), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	return e.Flush()
}

type Parent struct {
	GroupID      string `xml:"groupId,omitempty"`
	ArtifactID   string `xml:"artifactId,omitempty"`
	Version      string `xml:"version,omitempty"`
	RelativePath string `xml:"relativePath,omitempty"`
}

type Organization struct {
	Name string `xml:"name,omitempty"`
	URL  string `xml:"url,omitempty"`
}

type License struct {
	Name         string `xml:"name,omitempty"`
	URL          string `xml:"url,omitempty"`
	Distribution string `xml:"distribution,omitempty"`
	Comments     string `xml:"comments,omitempty"`
}

type Developer struct {
	ID              string      `xml:"id,omitempty"`
	Name            string      `xml:"name,omitempty"`
	Email           string      `xml:"email,omitempty"`
	URL             string      `xml:"url,omitempty"`
	Organization    string      `xml:"organization,omitempty"`
	OrganizationURL string      `xml:"organizationUrl,omitempty"`
	Roles           *[]string   `xml:"roles>role,omitempty"`
	Timezone        string      `xml:"timezone,omitempty"`
	Properties      *Properties `xml:"properties,omitempty"`
}

type Contributor struct {
	Name            string      `xml:"name,omitempty"`
	Email           string      `xml:"email,omitempty"`
	URL             string      `xml:"url,omitempty"`
	Organization    string      `xml:"organization,omitempty"`
	OrganizationURL string      `xml:"organizationUrl,omitempty"`
	Roles           *[]string   `xml:"roles>role,omitempty"`
	Timezone        string      `xml:"timezone,omitempty"`
	Properties      *Properties `xml:"properties,omitempty"`
}

type MailingList struct {
	Name          string    `xml:"name,omitempty"`
	Subscribe     string    `xml:"subscribe,omitempty"`
	Unsubscribe   string    `xml:"unsubscribe,omitempty"`
	Post          string    `xml:"post,omitempty"`
	Archive       string    `xml:"archive,omitempty"`
	OtherArchives *[]string `xml:"otherArchives>otherArchive,omitempty"`
}

type Prerequisites struct {
	Maven string `xml:"maven,omitempty"`
}

type Scm struct {
	Connection          string `xml:"connection,omitempty"`
	DeveloperConnection string `xml:"developerConnection,omitempty"`
	Tag                 string `xml:"tag,omitempty"`
	URL                 string `xml:"url,omitempty"`
}

type IssueManagement struct {
	System string `xml:"system,omitempty"`
	URL    string `xml:"url,omitempty"`
}

type CIManagement struct {
	System    string      `xml:"system,omitempty"`
	URL       string      `xml:"url,omitempty"`
	Notifiers *[]Notifier `xml:"notifiers>notifier,omitempty"`
}

type Notifier struct {
	Type          string      `xml:"type,omitempty"`
	SendOnError   bool        `xml:"sendOnError,omitempty"`
	SendOnFailure bool        `xml:"sendOnFailure,omitempty"`
	SendOnSuccess bool        `xml:"sendOnSuccess,omitempty"`
	SendOnWarning bool        `xml:"sendOnWarning,omitempty"`
	Address       string      `xml:"address,omitempty"`
	Configuration *Properties `xml:"configuration,omitempty"`
}

type DistributionManagement struct {
	Repository         *Repository `xml:"repository,omitempty"`
	SnapshotRepository *Repository `xml:"snapshotRepository,omitempty"`
	Site               *Site       `xml:"site,omitempty"`
	DownloadURL        string      `xml:"downloadUrl,omitempty"`
	Relocation         *Relocation `xml:"relocation,omitempty"`
	Status             string      `xml:"status,omitempty"`
}

type Site struct {
	ID   string `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	URL  string `xml:"url,omitempty"`
}

type Relocation struct {
	GroupID    string `xml:"groupId,omitempty"`
	ArtifactID string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
	Message    string `xml:"message,omitempty"`
}

type DependencyManagement struct {
	Dependencies *[]Dependency `xml:"dependencies>dependency,omitempty"`
}

type Dependency struct {
	GroupID    string       `xml:"groupId,omitempty"`
	ArtifactID string       `xml:"artifactId,omitempty"`
	Version    string       `xml:"version,omitempty"`
	Type       string       `xml:"type,omitempty"`
	Classifier string       `xml:"classifier,omitempty"`
	Scope      string       `xml:"scope,omitempty"`
	SystemPath string       `xml:"systemPath,omitempty"`
	Exclusions *[]Exclusion `xml:"exclusions>exclusion,omitempty"`
	Optional   string       `xml:"optional,omitempty"`
}

type Exclusion struct {
	GroupID    string `xml:"groupId,omitempty"`
	ArtifactID string `xml:"artifactId,omitempty"`
}

type Repository struct {
	UniqueVersion bool              `xml:"uniqueVersion,omitempty"`
	Releases      *RepositoryPolicy `xml:"releases,omitempty"`
	Snapshots     *RepositoryPolicy `xml:"snapshots,omitempty"`
	ID            string            `xml:"id,omitempty"`
	Name          string            `xml:"name,omitempty"`
	URL           string            `xml:"url,omitempty"`
	Layout        string            `xml:"layout,omitempty"`
}

type RepositoryPolicy struct {
	Enabled        string `xml:"enabled,omitempty"`
	UpdatePolicy   string `xml:"updatePolicy,omitempty"`
	ChecksumPolicy string `xml:"checksumPolicy,omitempty"`
}

type PluginRepository struct {
	Releases  *RepositoryPolicy `xml:"releases,omitempty"`
	Snapshots *RepositoryPolicy `xml:"snapshots,omitempty"`
	ID        string            `xml:"id,omitempty"`
	Name      string            `xml:"name,omitempty"`
	URL       string            `xml:"url,omitempty"`
	Layout    string            `xml:"layout,omitempty"`
}

type BuildBase struct {
	DefaultGoal      string            `xml:"defaultGoal,omitempty"`
	Resources        *[]Resource       `xml:"resources>resource,omitempty"`
	TestResources    *[]Resource       `xml:"testResources>testResource,omitempty"`
	Directory        string            `xml:"directory,omitempty"`
	FinalName        string            `xml:"finalName,omitempty"`
	Filters          *[]string         `xml:"filters>filter,omitempty"`
	PluginManagement *PluginManagement `xml:"pluginManagement,omitempty"`
	Plugins          *[]Plugin         `xml:"plugins>plugin,omitempty"`
}

type Build struct {
	SourceDirectory       string       `xml:"sourceDirectory,omitempty"`
	ScriptSourceDirectory string       `xml:"scriptSourceDirectory,omitempty"`
	TestSourceDirectory   string       `xml:"testSourceDirectory,omitempty"`
	OutputDirectory       string       `xml:"outputDirectory,omitempty"`
	TestOutputDirectory   string       `xml:"testOutputDirectory,omitempty"`
	Extensions            *[]Extension `xml:"extensions>extension,omitempty"`
	BuildBase
}

type Extension struct {
	GroupID    string `xml:"groupId,omitempty"`
	ArtifactID string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
}

type Resource struct {
	TargetPath string    `xml:"targetPath,omitempty"`
	Filtering  string    `xml:"filtering,omitempty"`
	Directory  string    `xml:"directory,omitempty"`
	Includes   *[]string `xml:"includes>include,omitempty"`
	Excludes   *[]string `xml:"excludes>exclude,omitempty"`
}

type PluginManagement struct {
	Plugins *[]Plugin `xml:"plugins>plugin,omitempty"`
}

// Configuration is a raw XML configuration that we currently do not muck with.
// It's supposed to be a DOM object, but upstream uses map[string]string for
// properties, and it does not work. For now, just keep it as a string so we
// can marshal it out untouched.
// TODO: This should be a DOM object.
// TODO: For some reason this loses the following XML attributes (probably all
// attributes):
// <configuration combine.children="merge">
// When marshalled, it is just:
// <configuration>
type Configuration struct {
	RawConfiguration string `xml:",innerxml"`
}

type Plugin struct {
	GroupID       string             `xml:"groupId,omitempty"`
	ArtifactID    string             `xml:"artifactId,omitempty"`
	Version       string             `xml:"version,omitempty"`
	Extensions    string             `xml:"extensions,omitempty"`
	Executions    *[]PluginExecution `xml:"executions>execution,omitempty"`
	Dependencies  *[]Dependency      `xml:"dependencies>dependency,omitempty"`
	Inherited     string             `xml:"inherited,omitempty"`
	Configuration *Configuration     `xml:"configuration,omitempty"`
}

type PluginExecution struct {
	ID        string    `xml:"id,omitempty"`
	Phase     string    `xml:"phase,omitempty"`
	Goals     *[]string `xml:"goals>goal,omitempty"`
	Inherited string    `xml:"inherited,omitempty"`
}

type Reporting struct {
	ExcludeDefaults string             `xml:"excludeDefaults,omitempty"`
	OutputDirectory string             `xml:"outputDirectory,omitempty"`
	Plugins         *[]ReportingPlugin `xml:"plugins>plugin,omitempty"`
}

type ReportingPlugin struct {
	GroupID    string       `xml:"groupId,omitempty"`
	ArtifactID string       `xml:"artifactId,omitempty"`
	Version    string       `xml:"version,omitempty"`
	Inherited  string       `xml:"inherited,omitempty"`
	ReportSets *[]ReportSet `xml:"reportSets>reportSet,omitempty"`
}

type ReportSet struct {
	ID        string    `xml:"id,omitempty"`
	Reports   *[]string `xml:"reports>report,omitempty"`
	Inherited string    `xml:"inherited,omitempty"`
}

type Profile struct {
	ID                     string                  `xml:"id,omitempty"`
	Activation             *Activation             `xml:"activation,omitempty"`
	Build                  *BuildBase              `xml:"build,omitempty"`
	Modules                *[]string               `xml:"modules>module,omitempty"`
	DistributionManagement *DistributionManagement `xml:"distributionManagement,omitempty"`
	Properties             *Properties             `xml:"properties,omitempty"`
	DependencyManagement   *DependencyManagement   `xml:"dependencyManagement,omitempty"`
	Dependencies           *[]Dependency           `xml:"dependencies>dependency,omitempty"`
	Repositories           *[]Repository           `xml:"repositories>repository,omitempty"`
	PluginRepositories     *[]PluginRepository     `xml:"pluginRepositories>pluginRepository,omitempty"`
	Reporting              *Reporting              `xml:"reporting,omitempty"`
}

type Activation struct {
	ActiveByDefault bool                `xml:"activeByDefault,omitempty"`
	JDK             string              `xml:"jdk,omitempty"`
	OS              *ActivationOS       `xml:"os,omitempty"`
	Property        *ActivationProperty `xml:"property,omitempty"`
	File            *ActivationFile     `xml:"file,omitempty"`
}

type ActivationOS struct {
	Name    string `xml:"name,omitempty"`
	Family  string `xml:"family,omitempty"`
	Arch    string `xml:"arch,omitempty"`
	Version string `xml:"version,omitempty"`
}

type ActivationProperty struct {
	Name  string `xml:"name,omitempty"`
	Value string `xml:"value,omitempty"`
}

type ActivationFile struct {
	Missing string `xml:"missing,omitempty"`
	Exists  string `xml:"exists,omitempty"`
}
