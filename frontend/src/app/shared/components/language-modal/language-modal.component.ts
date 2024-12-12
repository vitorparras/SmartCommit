import { Component } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { TranslateService } from '@ngx-translate/core';

interface Language {
  code: string;
  name: string;
  nativeName: string;
  englishName: string;
}

@Component({
  selector: 'app-language-modal',
  templateUrl: './language-modal.component.html',
  styleUrls: ['./language-modal.component.scss']
})
export class LanguageModalComponent {
  languages: Language[] = [
    { code: 'us', name: 'English', nativeName: 'English', englishName: 'English' },
    { code: 'br', name: 'Português', nativeName: 'Português', englishName: 'Portuguese' },
    { code: 'es', name: 'Español', nativeName: 'Español', englishName: 'Spanish' },
    { code: 'fr', name: 'Français', nativeName: 'Français', englishName: 'French' },
    { code: 'de', name: 'Deutsch', nativeName: 'Deutsch', englishName: 'German' }
  ];

  constructor(
    public dialogRef: MatDialogRef<LanguageModalComponent>,
    public translate: TranslateService
  ) {}

  selectLanguage(langCode: string): void {
    this.translate.use(langCode);
    this.dialogRef.close(langCode);
  }

  handleImageError(event: Event): void {
    const img = event.target as HTMLImageElement;
    img.style.display = 'none';
  }

  isCurrentLanguage(langCode: string): boolean {
    return this.translate.currentLang === langCode || 
           (!this.translate.currentLang && langCode === 'us');
  }
}

