			_, err := io.Copy(io.Discard, reader)
			workdir, err := os.MkdirTemp("", "empty-work-dir")