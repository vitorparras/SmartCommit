import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';

import { AIModel, CommitConfig, CommitFormat, CommitResult, CommitStyle, Language } from 'src/app/core/models/commit.model';
import { AIModelService } from 'src/app/core/services/ai-model.service';
import { CommitService } from 'src/app/core/services/commit.service';
import { RepositoryService } from 'src/app/core/services/repository.service';
import { ApiKeyModalComponent } from '../api-key-modal/api-key-modal.component';
import { ApiKeyService } from 'src/app/core/services/api-key.service';

@Component({
  selector: 'app-commit-generator',
  templateUrl: './commit-generator.component.html',
  styleUrls: ['./commit-generator.component.scss']
})
export class CommitGeneratorComponent implements OnInit {
  commitForm: FormGroup;
  configExpanded = true;
  isLoading = false;
  commitResult: any = null;

  aiModels: AIModel[] = [
    { id: 'gpt3', name: 'GPT-3', isOnline: true },
    { id: 'gpt4', name: 'GPT-4', isOnline: true },
    { id: 'codex', name: 'Codex', isOnline: true },
    { id: 'local', name: 'Local AI', isOnline: false }
  ];

  constructor(
    private fb: FormBuilder,
    private commitService: CommitService,
    private apiKeyService: ApiKeyService,
    private dialog: MatDialog,
    private snackBar: MatSnackBar
  ) {
    this.commitForm = this.fb.group({
      format: ['Angular Guidelines', Validators.required],
      repoPath: ['', Validators.required],
      aiModel: ['gpt4', Validators.required],
      style: ['Descriptive', Validators.required],
      language: ['English', Validators.required]
    });

    this.commitForm.get('aiModel')?.valueChanges.subscribe(async (modelId) => {
      const model = this.aiModels.find(m => m.id === modelId);
      if (model?.isOnline && !this.apiKeyService.getApiKey(modelId)) {
        const apiKey = await this.openApiKeyModal();
        if (apiKey) {
          this.apiKeyService.saveApiKey(modelId, apiKey);
        }
      }
    });
  }

  ngOnInit() {}

  async selectFolder() {
    try {
      //const dirHandle = await window.showDirectoryPicker();
      const gitPath ="testess" //await this.findGitRepository(dirHandle);
      
      if (gitPath) {
        this.commitForm.patchValue({ repoPath: gitPath });
      } else {
        this.snackBar.open('Selected folder is not a Git repository', 'Close', {
          duration: 3000
        });
      }
    } catch (error) {
      console.error('Error selecting folder:', error);
      this.snackBar.open('Failed to select folder', 'Close', {
        duration: 3000
      });
    }
  }

  private async findGitRepository(dirHandle: FileSystemDirectoryHandle): Promise<string | null> {
    try {
      await dirHandle.getDirectoryHandle('.git');
      return dirHandle.name;
    } catch {
      // If .git folder is not found, return null
      return null;
    }
  }

  async onSubmit() {
    if (this.commitForm.valid) {
      this.isLoading = true;
      try {
        this.commitResult = await this.commitService.generateCommitMessage(this.commitForm.value);
      } catch (error) {
        this.snackBar.open('Error generating commit message', 'Close', {
          duration: 3000
        });
      } finally {
        this.isLoading = false;
      }
    }
  }

  private openApiKeyModal() {
    const dialogRef = this.dialog.open(ApiKeyModalComponent, {
      width: '400px',
      disableClose: true
    });
    return dialogRef.afterClosed().toPromise();
  }

  copyToClipboard(text: string) {
    navigator.clipboard.writeText(text).then(() => {
      this.snackBar.open('Copied to clipboard', 'Close', {
        duration: 2000
      });
    });
  }
}