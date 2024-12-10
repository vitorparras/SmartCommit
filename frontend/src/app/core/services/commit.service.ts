import { Injectable } from '@angular/core';
import { CommitConfig, CommitResult } from '../models/commit.model';

@Injectable({
  providedIn: 'root'
})
export class CommitService {
  constructor() {}

  async generateCommitMessage(config: CommitConfig): Promise<CommitResult> {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));

    // Mock response
    const prefix = config.format === 'Conventional Commits' ? 'feat: ' : '';
    const style = config.style === 'Descriptive' ? 'Implement new feature for improved user experience' : 'Add new feature';
    const englishMessage = `${prefix}${style} using ${config.aiModel}`;
    
    const translatedMessage = config.language === 'English' 
      ? englishMessage 
      : `${prefix}Translated: ${style} using ${config.aiModel} (to ${config.language})`;

    return { englishMessage, translatedMessage };
  }
}

