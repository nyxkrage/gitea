	"code.gitea.io/gitea/models/db"
	"code.gitea.io/gitea/modules/analyze"
	"code.gitea.io/gitea/modules/util"
	IsGenerated             bool
	IsVendored              bool
	Start, End                             string

			lastFile := createDiffFile(diff, line)
			diff.End = lastFile.Name
			_, err := io.Copy(io.Discard, reader)
				count, err := db.Count(m)
func GetDiffRangeWithWhitespaceBehavior(gitRepo *git.Repository, beforeCommitID, afterCommitID, skipTo string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string, directComparison bool) (*Diff, error) {
	argsLength := 6
	if len(whitespaceBehavior) > 0 {
		argsLength++
	}
	if len(skipTo) > 0 {
		argsLength++
	}

	diffArgs := make([]string, 0, argsLength)
		diffArgs = append(diffArgs, "diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M")
		diffArgs = append(diffArgs, "diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M")
	if skipTo != "" {
		diffArgs = append(diffArgs, "--skip-to="+skipTo)
	}
	cmd := exec.CommandContext(ctx, git.GitExecutable, diffArgs...)

	diff.Start = skipTo

	var checker *git.CheckAttributeReader

	if git.CheckGitVersionAtLeast("1.7.8") == nil {
		indexFilename, deleteTemporaryFile, err := gitRepo.ReadTreeToTemporaryIndex(afterCommitID)
		if err == nil {
			defer deleteTemporaryFile()
			workdir, err := os.MkdirTemp("", "empty-work-dir")
			if err != nil {
				log.Error("Unable to create temporary directory: %v", err)
				return nil, err
			}
			defer func() {
				_ = util.RemoveAll(workdir)
			}()

			checker = &git.CheckAttributeReader{
				Attributes: []string{"linguist-vendored", "linguist-generated"},
				Repo:       gitRepo,
				IndexFile:  indexFilename,
				WorkTree:   workdir,
			}
			ctx, cancel := context.WithCancel(git.DefaultContext)
			if err := checker.Init(ctx); err != nil {
				log.Error("Unable to open checker for %s. Error: %v", afterCommitID, err)
			} else {
				go func() {
					err := checker.Run()
					if err != nil && err != ctx.Err() {
						log.Error("Unable to open checker for %s. Error: %v", afterCommitID, err)
					}
					cancel()
				}()
			}
			defer func() {
				cancel()
			}()
		}
	}


		gotVendor := false
		gotGenerated := false
		if checker != nil {
			attrs, err := checker.CheckPath(diffFile.Name)
			if err == nil {
				if vendored, has := attrs["linguist-vendored"]; has {
					if vendored == "set" || vendored == "true" {
						diffFile.IsVendored = true
						gotVendor = true
					} else {
						gotVendor = vendored == "false"
					}
				}
				if generated, has := attrs["linguist-generated"]; has {
					if generated == "set" || generated == "true" {
						diffFile.IsGenerated = true
						gotGenerated = true
					} else {
						gotGenerated = generated == "false"
					}
				}
			} else {
				log.Error("Unexpected error: %v", err)
			}
		}

		if !gotVendor {
			diffFile.IsVendored = analyze.IsVendor(diffFile.Name)
		}
		if !gotGenerated {
			diffFile.IsGenerated = analyze.IsGenerated(diffFile.Name)
		}

	separator := "..."
	if directComparison {
		separator = ".."
	}

	shortstatArgs := []string{beforeCommitID + separator + afterCommitID}
func GetDiffCommitWithWhitespaceBehavior(gitRepo *git.Repository, commitID, skipTo string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string, directComparison bool) (*Diff, error) {
	return GetDiffRangeWithWhitespaceBehavior(gitRepo, "", commitID, skipTo, maxLines, maxLineCharacters, maxFiles, whitespaceBehavior, directComparison)