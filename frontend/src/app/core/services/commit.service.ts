import { Injectable } from '@angular/core';
import { CommitConfig } from '../models/commit.model';

@Injectable({
  providedIn: 'root'
})
export class CommitService {
  async generateCommitMessage(config: CommitConfig) {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000));

    // Mock response
    return {
      englishMessage: `Add new feature using ${config.aiModel}`,
      translatedMessage: `Translated: Add new feature using ${config.aiModel} (to ${config.language})`
    };
  }
}

