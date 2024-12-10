import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';


import { AIModel, CommitConfig, CommitFormat, CommitResult, CommitStyle, Language } from 'src/app/core/models/commit.model';
import { AIModelService } from 'src/app/core/services/ai-model.service';
import { CommitService } from 'src/app/core/services/commit.service';
import { RepositoryService } from 'src/app/core/services/repository.service';
import { ApiKeyModalComponent } from '../api-key-modal/api-key-modal.component';

@Component({
  selector: 'app-commit-generator',
  templateUrl: './commit-generator.component.html',
  styleUrls: ['./commit-generator.component.css']
})
export class CommitGeneratorComponent implements OnInit {
  commitForm: FormGroup;
  aiModels: AIModel[] = [];
  languages: string[] = ['English', 'Spanish', 'French', 'German', 'Chinese'];
  commitFormats: string[] = ['Conventional Commits', 'Angular Guidelines', 'Custom'];
  commitStyles: string[] = ['Descriptive', 'Concise'];
  commitResult: CommitResult | null = null;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private commitService: CommitService,
    private aiModelService: AIModelService,
    private repositoryService: RepositoryService,
    private dialog: MatDialog,
    private snackBar: MatSnackBar
  ) {
    this.commitForm = this.fb.group({
      format: ['', Validators.required],
      repoPath: ['', Validators.required],
      aiModel: ['', Validators.required],
      style: ['', Validators.required],
      language: ['English', Validators.required]
    });
  }

  ngOnInit() {
    this.loadAIModels();
  }

  async loadAIModels() {
    try {
      this.aiModels = await this.aiModelService.getAvailableModels();
    } catch (error) {
      this.showErrorMessage('Failed to load AI models. Please try again.');
    }
  }

  async onSubmit() {
    if (this.commitForm.valid) {
      const config: CommitConfig = this.commitForm.value;
      const selectedModel = this.aiModels.find(model => model.id === config.aiModel);

      if (selectedModel && selectedModel.isOnline) {
        const apiKey = await this.openApiKeyModal();
        if (!apiKey) return;
        config.apiKey = apiKey;
      }

      this.isLoading = true;
      try {
        const isValidRepo = await this.repositoryService.validateRepository(config.repoPath);
        if (!isValidRepo) {
          throw new Error('Invalid repository path');
        }
        this.commitResult = await this.commitService.generateCommitMessage(config);
      } catch (error) {
        this.showErrorMessage('Error generating commit message. Please try again.');
      } finally {
        this.isLoading = false;
      }
    }
  }

  openApiKeyModal(): Promise<string | null> {
    const dialogRef = this.dialog.open(ApiKeyModalComponent);
    return dialogRef.afterClosed().toPromise();
  }

  private showErrorMessage(message: string) {
    this.snackBar.open(message, 'Close', {
      duration: 5000,
      horizontalPosition: 'center',
      verticalPosition: 'bottom',
    });
  }
}