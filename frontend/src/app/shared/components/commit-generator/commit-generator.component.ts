import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ApiKeyModalComponent } from '../api-key-modal/api-key-modal.component';
import { TranslateService } from '@ngx-translate/core';
import { CommitService } from 'src/app/core/services/commit.service';
import { ApiKeyService } from 'src/app/core/services/api-key.service';

interface AIModel {
  id: string;
  name: string;
  isOnline: boolean;
}

@Component({
  selector: 'app-commit-generator',
  templateUrl: './commit-generator.component.html',
  styleUrls: ['./commit-generator.component.scss']
})
export class CommitGeneratorComponent implements OnInit {
  @ViewChild('fileInput') fileInput!: ElementRef<HTMLInputElement>;
  
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

  languages = [
    { code: 'us', name: 'English' },
    { code: 'es', name: 'Español' },
    { code: 'fr', name: 'Français' },
    { code: 'de', name: 'Deutsch' },
    { code: 'br', name: 'Português (Brasil)' }
  ];

  constructor(
    private fb: FormBuilder,
    private commitService: CommitService,
    private apiKeyService: ApiKeyService,
    private dialog: MatDialog,
    private snackBar: MatSnackBar,
    private translate: TranslateService
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

  ngOnInit() {
    // Set initial language
    this.translate.setDefaultLang('us');
  }

  selectGitFolder() {
    this.fileInput.nativeElement.click();
  }

  handleFileSelection(event: Event) {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files[0]) {
      const file = input.files[0];
      if (file.name.toLowerCase() === 'readme.md') {
        this.commitForm.patchValue({ repoPath: this.translate.instant('SELECT_INDEX_FILE_SUCCESS') });
      } else {
        this.snackBar.open(this.translate.instant('SELECT_README_FILE'), this.translate.instant('CLOSE'), {
          duration: 3000
        });
      }
    }
  }

  async onSubmit() {
    if (this.commitForm.valid) {
      this.isLoading = true;
      this.configExpanded = false; // Collapse configuration section
      try {
        this.commitResult = await this.commitService.generateCommitMessage(this.commitForm.value);
      } catch (error) {
        this.snackBar.open(this.translate.instant('ERROR_GENERATING_COMMIT'), this.translate.instant('CLOSE'), {
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
      this.snackBar.open(this.translate.instant('COPIED_TO_CLIPBOARD'), this.translate.instant('CLOSE'), {
        duration: 2000
      });
    });
  }
}

