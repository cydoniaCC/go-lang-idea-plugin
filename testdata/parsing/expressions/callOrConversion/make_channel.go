package main
var e = make(chan int)
/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  VarDeclarationsImpl
    PsiElement(KEYWORD_VAR)('var')
    PsiWhiteSpace(' ')
    VarDeclarationImpl
      Identifiers
        PsiElement(IDENTIFIER)('e')
      PsiWhiteSpace(' ')
      PsiElement(=)('=')
      PsiWhiteSpace(' ')
      BuiltInCallExpressionImpl
        LiteralExpressionImpl
          PsiElement(IDENTIFIER)('make')
        PsiElement(()('(')
        TypeChanBidiImpl
          PsiElement(KEYWORD_CHAN)('chan')
          PsiWhiteSpace(' ')
          TypeNameImpl
            PsiElement(IDENTIFIER)('int')
        PsiElement())(')')
