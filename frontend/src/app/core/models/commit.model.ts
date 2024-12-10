export interface CommitFormat {
  value: string;
  label: string;
}

export interface CommitStyle {
  value: string;
  label: string;
}

export interface Language {
  value: string;
  label: string;
}

export interface CommitConfig {
  format: string;
  repoPath: string;
  aiModel: string;
  style: string;
  language: string;
  apiKey?: string;
}

export interface CommitResult {
  englishMessage: string;
  translatedMessage: string;
}

export interface AIModel {
  id: string;
  name: string;
  isOnline: boolean;
}



